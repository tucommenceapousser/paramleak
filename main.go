package main

import (
		"github.com/logrusorgru/aurora"
		// autres imports existants
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/tucommenceapousser/paramleak/regex"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	fUtils "github.com/projectdiscovery/utils/file"
	httpUtil "github.com/projectdiscovery/utils/http"
	sUtils "github.com/projectdiscovery/utils/slice"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type options struct {
	url     string
	list    string
	delay   time.Duration
	silent  bool
	method  string
	body    string
	header  string
	verbose bool
}

var (
	opt = &options{}
)

func main() {
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`Paramleak is a tool for get all parameters from input`)
	flagSet.StringVarP(&opt.url, "url", "u", "", "url for get parameters")
	flagSet.StringVarP(&opt.list, "list", "l", "", "List of Url for Get All parameters")
	flagSet.StringVarP(&opt.method, "method", "X", "GET", "Http Method for request")
	flagSet.StringVarP(&opt.body, "body", "d", "", "Body data for Post Request")
	flagSet.StringVarP(&opt.header, "header", "H", "", "Custom Header (You can set only one custom header)")
	flagSet.DurationVarP(&opt.delay, "delay", "p", 0, "time for delay example: 1000 Millisecond (1 second)")
	flagSet.BoolVarP(&opt.verbose, "verbose", "v", false, "Verbose mode")
	flagSet.BoolVarP(&opt.silent, "silent", "s", false, "Silent mode")

	if err := flagSet.Parse(); err != nil {
		gologger.Error().Msgf("Could not parse flags: %s", err)
	}

	if opt.url == "" && opt.list == "" && !fUtils.HasStdin() {
		printUsage()
		return
	}

	if !isValidHTTPMethod(opt.method) {
		gologger.Error().Msgf("Invalid HTTP method!\nAllowed methods: GET, POST, PUT, PATCH, DELETE")
		return
	}

	var parts = strings.SplitN(opt.header, ":", 2)
	if opt.header != "" && len(parts) != 2 {
		gologger.Error().Msgf("invalid custom header! (Your input must be two part separate with ':')\n\t -H 'Cookie: test=test;'")
		return
	}

	if !opt.silent {
		Banner()
	}

	input := getInput()
	if len(input) > 0 {
		results := run(input, opt.method, opt.body, opt.delay)
		for _, v := range results {
			fmt.Println(v)
		}
	}
}

func getInput() []string {
	input := make([]string, 0)
	switch {
	case fUtils.HasStdin():
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
	case opt.url != "":
		input = append(input, opt.url)
	case opt.list != "":
		lists, err := fUtils.ReadFile(opt.list)
		if err != nil {
			gologger.Error().Msgf("%v", err)
			return nil
		}
		for content := range lists {
			input = append(input, content)
		}
	}
	return input
}

func run(lists []string, method, body string, delay time.Duration) []string {
	output := make([]string, 0)
	c := make(chan []string)

	for _, Url := range lists {
		go func(URL string) {
			time.Sleep(delay * time.Millisecond)
			data, err := Request(method, URL, body)
			if err != nil {
				c <- []string{}
			}
			result := regex.Regex(data)
			c <- result
		}(Url)
	}

	for range lists {
		result := <-c
		for _, s := range result {
			reg := regexp.MustCompile(" ")
			s = reg.ReplaceAllString(s, "")
			output = append(output, s)
		}
	}

	return sUtils.Dedupe(output)
}

var client = &http.Client{
	Timeout: 3 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func Request(method, urlStr, bodyStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if u.Scheme == "" {
		urlStr = "http://" + u.Host + u.Path
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	reqBody := strings.NewReader(bodyStr)
	req, err := http.NewRequestWithContext(ctx, method, urlStr, reqBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/112.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Origin", "https://"+u.Host)

	var parts = strings.SplitN(opt.header, ":", 2)
	if opt.header != "" {
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Set custom header
		req.Header.Set(key, value)
	}

	if opt.verbose {
		dumpRequest, err := httpUtil.DumpRequest(req)
		if err != nil {
			return "", err
		}
		gologger.Print().Msgf(dumpRequest)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if opt.verbose {
		dumpHeaderResponse, dumpBodyResponse, err := httpUtil.DumpResponseHeadersAndRaw(resp)
		if err != nil {
			return "", err
		}

		gologger.Print().Msgf(string(dumpHeaderResponse))
		gologger.Print().Msgf(string(dumpBodyResponse))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func isValidHTTPMethod(method string) bool {
	validMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	for _, m := range validMethods {
		if method == m {
			return true
		}
	}
	return false
}

func printUsage() {
	Banner()
//	gologger.Print().Msgf("Flags:\n")
//	gologger.Print().Msgf("\t-url,      -u       URL for getting all parameters")
		fmt.Println(aurora.Bold(aurora.Green("Flags:\n")))

		fmt.Println(
			aurora.Bold(aurora.Red("\t-url, -u")).String() +
				aurora.Bold(aurora.Blue("       URL for getting all parameters")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-list, -l")).String() +
				aurora.Bold(aurora.Blue("       List of URLs for getting all parameters")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-method, -X")).String() +
				aurora.Bold(aurora.Blue("       HTTP method for requests")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-body, -d")).String() +
				aurora.Bold(aurora.Blue("       Body data for POST/PATCH requests")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-header, -H")).String() +
				aurora.Bold(aurora.Blue("       Custom header (You can set only 1 custom header)")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-delay, -p")).String() +
				aurora.Bold(aurora.Blue("       Delay time example: 1000 milliseconds (1 second)")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-verbose, -v")).String() +
				aurora.Bold(aurora.Blue("       Verbose mode")).String(),
		)

		fmt.Println(
			aurora.Bold(aurora.Red("\t-silent, -s")).String() +
				aurora.Bold(aurora.Blue("       Silent mode")).String(),
		)
	}

func printSuccess(msg string) {
	fmt.Println(aurora.Green(msg))
}

func printError(msg string) {
	fmt.Println(aurora.Red(msg))
}

func printInfo(msg string) {
	fmt.Println(aurora.Cyan(msg))
}

func Banner() {
	fmt.Println(aurora.Bold(aurora.Magenta(`
██████╗  █████╗ ██████╗  █████╗ ███╗   ███╗██╗     ███████╗ █████╗ ██╗  ██╗
██╔══██╗██╔══██╗██╔══██╗██╔══██╗████╗ ████║██║     ██╔════╝██╔══██╗██║ ██╔╝
██████╔╝███████║██████╔╝███████║██╔████╔██║██║     █████╗  ███████║█████╔╝ 
██╔═══╝ ██╔══██║██╔══██╗██╔══██║██║╚██╔╝██║██║     ██╔══╝  ██╔══██║██╔═██╗ 
██║     ██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝
	`)))
	fmt.Println(aurora.Bold(aurora.Cyan("\t\tModded by Trhacknon :)\n\n")))
}

