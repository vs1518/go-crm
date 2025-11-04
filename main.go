package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type contact struct {
	ID    int
	name  string
	email string
}

var contacts map[int]*contact

var (
	name  = flag.String("name", "", "Name of the contact")
	email = flag.String("email", "", "Email of the contact")
)

func newContact(contact *contact) {
	if contact.name == "" || contact.email == "" {
		fmt.Print("Incorrect: name or email cannot be empty\n")
		miniCRM()
	}
	
	// Validate email format with regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(contact.email) {
		fmt.Print("Incorrect: invalid email format\n")
		miniCRM()
	}
	
	fmt.Print("Contact created successfully\n")
}

func createId() int {
	// iterate over keys to get the latest id
	if len(contacts) == 0 {
		return 1
	}

	maxId := 0
	for _, contact := range contacts {
		if contact.ID > maxId {
			maxId = contact.ID
		}
	}

	return maxId + 1
}

func printItems() []string {
	choices := []string{
		"1) Add contact",
		"2) List contacts",
		"3) Remove contact",
		"4) Update contact",
		"5) Quit",
	}
	for i := 0; i < len(choices); i++ {
		fmt.Print(choices[i] + "\n")

	}
	return choices
}

func addContact(reader bufio.Reader) {
	// prefix string "Enter contact name: "
	id := createId()
	fmt.Print("Enter contact name: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Print("Enter contact email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	// append to contacts (indexed by name) a new contact with name input and empty string as value
	contact := &contact{ID: id, name: input, email: email}
	newContact(contact)
	contacts[id] = contact
	fmt.Print("Contact added: " + contact.name + "\n")
	miniCRM()
}

func ListContacts(reader bufio.Reader) {
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Email: %s, Name: %s\n", contact.ID, contact.email, contact.name)
	}
	fmt.Print("Return to menu ? (y/n)")
	input, _ := reader.ReadString('\n')
	switch {
	case input == "y\n":
		miniCRM()
	default:
		miniCRM()
	}
}

func updateContact(ID int, newName, newEmail string) {
	contact, exists := contacts[ID]
	if !exists {
		fmt.Printf("Contact not found: ID: %d\n", ID)
		return
	}

	// Store the old contact info for display
	oldName := contact.name
	oldEmail := contact.email

	// Update the contact
	contact.name = newName
	contact.email = newEmail

	fmt.Printf("Contact updated: ID: %d\n", contact.ID)
	fmt.Printf("Old Name: %sOld Email: %s\n", oldName, oldEmail)
	fmt.Printf("New Name: %sNew Email: %s\n", newName, newEmail)
}

func removeContact(ID int) {
	// iterate over contacts, find the one with corresponding ID, remove it
	contact, exists := contacts[ID]
	if !exists {
		fmt.Printf("Contact not found: ID: %d\n", ID)
		return
	}
	delete(contacts, ID)
	fmt.Printf("Contact removed: ID: %d, Email: %s, Name: %s\n", contact.ID, contact.email, contact.name)
}

func handleRemoveContact(reader bufio.Reader) {
	for {
		fmt.Print("Enter contact ID to remove: ")
		idInput, _ := reader.ReadString('\n')

		// Remove newline and parse as integer
		idStr := idInput[:len(idInput)-1]
		if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
			removeContact(id)
			break
		}
		fmt.Println("Invalid ID. Please enter a positive number.")
	}
	miniCRM()
}

func handleUpdateContact(reader bufio.Reader) {
	for {
		fmt.Print("Enter contact ID to update: ")
		idInput, _ := reader.ReadString('\n')

		// Remove newline and parse as integer
		idStr := idInput[:len(idInput)-1]
		if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
			// Check if contact exists
			if _, exists := contacts[id]; !exists {
				fmt.Printf("Contact not found: ID: %d\n", id)
				continue
			}

			// Display current contact info
			currentContact := contacts[id]
			fmt.Printf("Current contact:\n")
			fmt.Printf("ID: %d\n", currentContact.ID)
			fmt.Printf("Name: %s", currentContact.name)
			fmt.Printf("Email: %s", currentContact.email)

			// Ask for new information
			fmt.Print("Enter new contact name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)
			fmt.Print("Enter new contact email: ")
			newEmail, _ := reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)

			// Update the contact
			updateContact(id, newName, newEmail)
			break
		}
		fmt.Println("Invalid ID. Please enter a positive number.")
	}
	miniCRM()
}

func miniCRM() {
	printItems()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Step 1 : if input not integer between 1 and 5, ask again
	for {
		if input == "1\n" || input == "2\n" || input == "3\n" || input == "4\n" || input == "5\n" {
			break
		}
		fmt.Print("Invalid input. Please enter a number between 1 and 5: ")
		input, _ = reader.ReadString('\n')
	}

	switch {
	case input == "1\n":
		{
			addContact(*reader)
		}
	case input == "2\n":
		{
			ListContacts(*reader)
		}
	case input == "3\n":
		{
			handleRemoveContact(*reader)
		}
	case input == "4\n":
		{
			handleUpdateContact(*reader)
		}
	case input == "5\n":
		{
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func main() {
	contacts = make(map[int]*contact)
	flag.Parse()
	name := strings.TrimSpace(*name)
	email := strings.TrimSpace(*email)
	if name != "" && email != "" {
		contacts[1] = &contact{
			ID:    1,
			name:  name,
			email: email,
		}
	}
	miniCRM()
}
