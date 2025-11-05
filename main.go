package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	name   = flag.String("name", "", "Name of the contact")
	email  = flag.String("email", "", "Email of the contact")
	nextid = 0
)

func validateContact(contact *contact) error {
	if contact.name == "" || contact.email == "" {
		return errors.New("name and email cannot be empty")
	}

	// Validate email format with regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(contact.email) {
		return errors.New("invalid email format")
	}

	return nil
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

func addContact(reader bufio.Reader, store storer) {
	// prefix string "Enter contact name: "
	id := nextid + 1
	nextid = id
	fmt.Print("\n" + "Enter contact name: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Print("Enter contact email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	contact := &contact{ID: id,
		name:  input,
		email: email}
	if err := validateContact(contact); err != nil {
		fmt.Printf("Validation failed: %s\n", err.Error())
		core(store)
	}
	store.save(contact)
	fmt.Print("Contact added: " + contact.name + " with ID: " + strconv.Itoa(contact.ID) + "\n")
	core(store)
}

func ListContacts(reader bufio.Reader, store storer) {
	// display all contacts in store memory with ID, email and name in table format
	store.render()
	fmt.Println("Press enter to return to menu...")
	input, _ := reader.ReadString('\n')
	if input == "\n" {
		core(store)
	}
}

func handleRemoveContact(reader bufio.Reader, store storer) {
	for {
		fmt.Print("Enter contact ID to remove: ")
		idInput, _ := reader.ReadString('\n')
		// Remove newline and parse as integer
		idStr := idInput[:len(idInput)-1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID. Please enter a positive number.")
			continue
		}
		if id > 0 {
			if err := store.delete(id); err != nil {
				fmt.Println("Error deleting contact:", err)
				continue
			}
			fmt.Println("Contact deleted successfully")
			core(store)
		}
	}
}

func handleUpdateContact(reader bufio.Reader, store storer) {
	for {
		fmt.Print("Enter contact ID to update: ")
		idInput, _ := reader.ReadString('\n')

		// Remove newline and parse as integer
		idStr := idInput[:len(idInput)-1]
		if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
			// Check if contact exists and display current contact info
			fmt.Printf("Current contact:\n")
			if err := store.renderOne(id); err != nil {
				fmt.Printf("Contact not found: ID: %d\n", id)
				continue
			}

			// Ask for new information
			fmt.Print("Enter new contact name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)
			fmt.Print("Enter new contact email: ")
			newEmail, _ := reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)

			// Update the contact
			newContact := &contact{ID: id,
				name:  newName,
				email: newEmail}
			store.update(newContact)
			break
		}
		fmt.Println("Invalid ID. Please enter a positive number.")
	}
	core(store)
}

func core(store storer) {
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
			addContact(*reader, store)
		}
	case input == "2\n":
		{
			ListContacts(*reader, store)
		}
	case input == "3\n":
		{
			handleRemoveContact(*reader, store)
		}
	case input == "4\n":
		{
			handleUpdateContact(*reader, store)
		}
	case input == "5\n":
		{
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func main() {
	flag.Parse()
	store := &MemoryStore{
		contacts: make(map[int]*contact),
	}
	name := strings.TrimSpace(*name)
	email := strings.TrimSpace(*email)
	if name != "" && email != "" {
		id := nextid + 1
		nextid = id
		contact := &contact{ID: id,
			name:  name,
			email: email,
		}
		store.save(contact)
	}
	core(store)
}
