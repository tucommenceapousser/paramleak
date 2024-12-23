# ⚡️ **paramleaks** ⚡️  
### A tool for extracting parameters from Websites (HTML/JSON)  
![GitHub Repo stars](https://img.shields.io/github/stars/tucommenceapousser/paramleak?style=for-the-badge&logo=github)

 
![anonymous](https://img.shields.io/badge/anonymous-000000?style=for-the-badge&logo=ghost&logoColor=FF0000)

![TRHACKNON](https://img.shields.io/badge/TRHACKNON-000000?style=for-the-badge&logo=three.js&logoColor=FF0000)

![tucommenceapousser](https://img.shields.io/badge/tucommenceapousser-000000?style=for-the-badge&logo=github&logoColor=FF0000)


![Kali](https://img.shields.io/badge/Kali-000000?style=for-the-badge&logo=kalilinux&logoColor=FF0000)

![Modded By](https://img.shields.io/badge/Modded%20by-Trhacknon-ff69b4?style=for-the-badge&logo=github)

![Followers](https://img.shields.io/github/followers/tucommenceapousser?style=for-the-badge&color=ff0000)

![Trhacknon's GitHub Stats](https://github-readme-stats.vercel.app/api?username=tucommenceapousser&show_icons=true&count_private=true&hide=prs&theme=tokyonight&bg_color=000000&title_color=ff0000&text_color=ff0000&layout=compact&border_color=0099ff)

![Trhacknon's Top Languages](https://github-readme-stats.vercel.app/api/top-langs/?username=tucommenceapousser&layout=compact&theme=highcontrast&border_color=1b03a3)

+ 🎮 **Mod'z from [SharokhAtaie](https://github.com/SharokhAtaie/paramleak)** 🎮


---

### Progression du projet
- [x] Extraction des paramètres depuis une URL
- [x] Extraction des paramètres depuis une liste d'URLs
- [x] Prise en charge de plusieurs méthodes HTTP
- [ ] Interface utilisateur graphique (en cours) ⏳
---

# Aperçu du fonctionnement

<details>
  <summary><strong>🔍 Menu 📸</strong></summary>
<p align="center">
  <img src="https://j.top4top.io/p_3261bypwh0.jpg" width="80%" />
</p>

</details>

---

<details>
  <summary><strong>🔍 Exécution avec une liste d'URLs 📸 
    &
    🔍 Exécution avec une URL 📸</strong></summary>

<p align="center">
  <img src="https://b.top4top.io/p_3261klgbn0.jpg" width="45%" />
  <img src="https://k.top4top.io/p_3261i55q31.jpg" width="45%" />
</p>
</details>

---

## 🚀 **Installation**  
<details>
<summary><strong>Cliquez pour afficher les instructions</strong></summary>

```diff
+ obligatoire
- eventuel
```
```diff
# Installer depuis la source
- go install github.com/tucommenceapousser/paramleak@latest
```
```diff
- git clone https://github.com/tucommenceapousser/paramleak
- go install
# Compiler
- go build -o paramleak .
+ chmod +x paramleak
```

</details>

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
<summary><strong>Cliquez pour afficher un exemple</strong></summary>
  
```ruby
# Lancer le scan avec une URL
./paramleak -u https://www.igrocity.ru/search.php
```
```ruby
# Lancer le scan avec une liste d'URLs
./paramleak -l list.txt
```
Résultat :

```ruby
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
```
</details>
---

💬 Remerciements

Un grand merci à SharokhAtaie pour le projet original ! 🙏


---

👿 Suivez le projet ou contribuez si vous le souhaitez, modifications faites par Trhacknon ! 😈
