# ⚡️ **paramleaks** ⚡️  
### A tool for extracting parameters from Websites (HTML/JSON)  
**Modded by [Trhacknon](#)**

---

## 🚀 **Installation**  
<details>
<summary><strong>Cliquez pour afficher les instructions</strong></summary>

```bash
# Installer depuis la source
go install github.com/tucommenceapousser/paramleak@latest
# Compiler
go build -o paramleak .
chmod +x paramleak
# Lancer le scan avec une liste d'URLs
./paramleak -l list.txt
```

</details>
---

🛠️ Usage

██████╗  █████╗ ██████╗  █████╗ ███╗   ███╗██╗     ███████╗ █████╗ ██╗  ██╗
██╔══██╗██╔══██╗██╔══██╗██╔══██╗████╗ ████║██║     ██╔════╝██╔══██╗██║ ██╔╝
██████╔╝███████║██████╔╝███████║██╔████╔██║██║     █████╗  ███████║█████╔╝ 
██╔═══╝ ██╔══██║██╔══██╗██╔══██║██║╚██╔╝██║██║     ██╔══╝  ██╔══██║██╔═██╗ 
██║     ██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝
                   Modded by Trhacknon 🦾


---

⚙️ Flags disponibles

<details>
<summary><strong>Cliquez pour voir tous les flags</strong></summary>

| Flag        | Alias | Description                                                 |
|-------------|-------|-------------------------------------------------------------|
| `-url`      | `-u`  | URL cible pour extraire les paramètres                      |
| `-list`     | `-l`  | Liste d'URLs pour extraire les paramètres                   |
| `-method`   | `-X`  | Méthode HTTP (GET, POST, PATCH, DELETE, PUT)                |
| `-body`     | `-d`  | Données de la requête pour POST/PATCH/DELETE/PUT            |
| `-header`   | `-H`  | Header personnalisé (un seul supporté)                      |
| `-delay`    | `-p`  | Délai entre les requêtes (en millisecondes, ex : 1000 = 1s) |
| `-verbose`  | `-v`  | Mode verbeux                                                |
| `-silent`   | `-s`  | Mode silencieux                                             |

</details>
---

🧪 Exemple d'utilisation

<details>
<summary><strong>Cliquez pour afficher un exemple</strong></summary>./paramleak -l list.txt

Résultat :

something
test_var
user_id_i
name
rdt_to
obj_key1
val1
obj_key2
val2
test_obj
empty_var
param1
method
param2

</details>
---

💬 Remerciements

Un grand merci à SharokhAtaie pour le projet original ! 🙏


---

🎉 Suivez le projet pour plus de modifications fun par Trhacknon ! 🎉
