package main

import (
  "bufio"  // lire efficacement depuis stdin
  "flag"   // gérer les arguments de la CLI
  "fmt"    // print/formatage
  "log"    // logs/erreurs fatales
  "os"     // accès système (stdin)
  "strconv"// conversion string<->int
  "strings"// trim/contains…
)
// Contact = entité métier minimale
type Contact struct {
	ID    int
	Name  string
	Email string
}

// "Base" en mémoire : map[ID]Contact
var contacts = make(map[int]Contact)

func main() {
	// --- Mode FLAGS : ajout direct sans passer par le menu ---
	addFlag := flag.Bool("add", false, "Add a contact via flags and exit")
	idFlag := flag.Int("id", 0, "Contact ID (int)")
	nameFlag := flag.String("name", "", "Contact name")
	emailFlag := flag.String("email", "", "Contact email")
	flag.Parse()

	if *addFlag {
		if err := addByFlags(*idFlag, *nameFlag, *emailFlag); err != nil {
			log.Fatalf("Failed to add contact by flags: %v", err)
		}
		fmt.Println("Contact added via flags ✅")
		return
	}

	// --- Mode INTERACTIF : menu en boucle ---
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		choice, err := readLine(reader, "Choix: ")
		if err != nil {
			fmt.Println("Erreur de lecture:", err)
			continue
		}
		switch strings.TrimSpace(choice) {
		case "1":
			if err := addInteractive(reader); err != nil {
				fmt.Println("❌", err)
			}
		case "2":
			listContacts()
		case "3":
			if err := deleteByID(reader); err != nil {
				fmt.Println("❌", err)
			}
		case "4":
			if err := updateContact(reader); err != nil {
				fmt.Println("❌", err)
			}
		case "5":
			fmt.Println("👋 Au revoir !")
			return
		default:
			fmt.Println("Option inconnue. Réessaie.")
		}
	}
}


