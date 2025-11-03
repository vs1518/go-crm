#  Mini-CRM – Application en Go

[![Langage](https://img.shields.io/badge/langage-Go-blue.svg?logo=go)](https://go.dev/)
[![Version](https://img.shields.io/badge/version-1.25.3-brightgreen)](https://go.dev/dl/)
[![Licence](https://img.shields.io/badge/Licence-MIT-lightgrey)](LICENSE)
[![Plateforme](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20MacOS-yellow.svg)](#)
[![Status](https://img.shields.io/badge/status-Actif-success)](#)

> **Projet pédagogique – EFREI Paris 2025–2026**
> Langage : **Go (Golang)**
> Auteur : **Arnaud D.* et *Valeriya S.*
> Encadrante : *Axelle Lança*

---

##  Objectif du projet

Ce projet consiste à créer une **application en ligne de commande (CLI)** simulant un **Mini CRM** (*Customer Relationship Management*).
L’objectif est de manipuler les **concepts fondamentaux du langage Go** à travers un cas concret : la gestion de contacts (ajout, suppression, mise à jour, affichage).

---

##  Fonctionnalités principales

| Fonctionnalité              | Description                                            |
| --------------------------- | ------------------------------------------------------ |
| ➕ Ajouter un contact        | Enregistre un contact avec ID, nom et email            |
| 📋 Lister les contacts      | Affiche tous les contacts enregistrés                  |
| 🗑️ Supprimer un contact    | Supprime un contact à partir de son ID                 |
| ✏️ Mettre à jour un contact | Modifie le nom ou l’email d’un contact                 |
| ⚙️ Flags CLI                | Permet l’ajout rapide via `--add --id --name --email`  |
| 💬 Interface interactive    | Menu en boucle avec saisie clavier (`bufio.NewReader`) |

---

##  Structure du projet

```
go-crm/
└── main.go             # Point d'entrée principal du programme
```

---

## 💻 Installation et exécution

### 1️⃣ Cloner le dépôt

```bash
git clone https://github.com/vs1518/go-crm.git
cd go-crm
```

### 2️⃣ Lancer le projet

```bash
go run main.go
```

Le menu s’affiche :

```
===== Mini-CRM =====
1) Ajouter un contact
2) Lister les contacts
3) Supprimer un contact
4) Mettre à jour un contact
5) Quitter
```

### 3️⃣ Ajouter via les flags (mode rapide)

```bash
go run main.go --add --id=1 --name="arno" --email="arno@gmail.com"
```

Résultat :

```
Contact added via flags ✅
```

---

##  Exemple d’utilisation

```
===== Mini-CRM =====
1) Ajouter un contact
2) Lister les contacts
3) Supprimer un contact par ID
4) Mettre à jour un contact
5) Quitter
Choix: 1
ID (entier > 0): 1
Nom: arno
Email: arno@gmail.com
✅ Contact ajouté.
```

---

##  Concepts Go utilisés

| Concept Go                  | Description                      |
| --------------------------- | -------------------------------- |
| `for {}`                    | Boucle infinie du menu principal |
| `switch`                    | Gestion du choix utilisateur     |
| `map[int]Contact`           | Stockage des contacts            |
| “comma ok idiom”            | Vérifie l’existence d’un contact |
| `if err != nil`             | Gestion d’erreurs                |
| `strconv.Atoi()`            | Conversion string → int          |
| `bufio.NewReader(os.Stdin)` | Lecture des entrées utilisateur  |
| `flag`                      | Gestion des arguments CLI        |

---

##  Technologies

| Type                  | Outil                     |
| --------------------- | ------------------------- |
| 💻 Langage            | Go 1.25.3                 |
| 🧰 IDE recommandés    | VS Code, GoLand, IntelliJ |
| 🔧 Gestion de modules | Go Modules (`go.mod`)     |

---

## 🔮 Améliorations possibles


* 🧪 Ajout de tests unitaires

---

##  Commandes Git utiles

| Action                       | Commande                                        |
| ---------------------------- | ----------------------------------------------- |
| Mettre à jour le dépôt local | `git pull origin main`                          |
| Ajouter les changements      | `git add .`                                     |
| Committer avec message       | `git commit -m "Ajout nouvelle fonctionnalité"` |
| Envoyer sur GitHub           | `git push origin main`                          |

---

##  Exemple de workflow complet

```bash
# 1. Mettre à jour depuis GitHub
git pull origin main

# 2. Modifier le code
code main.go  # ou via VS Code

# 3. Tester localement
go run main.go

# 4. Sauvegarder et pousser
git add .
git commit -m "Ajout fonction de mise à jour de contact"
git push origin main
```

---

## 📚 Contexte académique

Ce projet a été réalisé dans le cadre du **module Go – Les Fondamentaux**, enseigné à l’**EFREI Paris** (année 2025–2026).
