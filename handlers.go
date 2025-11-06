package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"minicrm/contact"
)

func validateContact(contact *contact.Contact) error {
	if contact.Name == "" || contact.Email == "" {
		return errors.New("name and email cannot be empty")
	}

	// Validate email format with regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(contact.Email) {
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

func addContact(reader bufio.Reader, Store contact.Storer) {
	// prefix string "Enter contact name: "
	id := nextid + 1
	nextid = id
	fmt.Print("\n" + "Enter contact name: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Print("Enter contact email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	contact := &contact.Contact{ID: id,
		Name:  input,
		Email: email}
	if err := validateContact(contact); err != nil {
		fmt.Printf("Validation failed: %s\n", err.Error())
		core(Store)
	}
	Store.Save(contact)
	fmt.Print("Contact added: " + contact.Name + " with ID: " + strconv.Itoa(contact.ID) + "\n")
	core(Store)
}

func ListContacts(reader bufio.Reader, Store contact.Storer) {
	// display all contacts in store memory with ID, email and name in table format
	Store.Render()
	fmt.Println("Press enter to return to menu...")
	input, _ := reader.ReadString('\n')
	if input == "\n" {
		core(Store)
	}
}

func handleRemoveContact(reader bufio.Reader, Store contact.Storer) {
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
			if err := Store.Delete(id); err != nil {
				fmt.Println("Error deleting contact:", err)
				continue
			}
			fmt.Println("Contact deleted successfully")
			core(Store)
		}
	}
}

func handleUpdateContact(reader bufio.Reader, Store contact.Storer) {
	for {
		fmt.Print("Enter contact ID to update: ")
		idInput, _ := reader.ReadString('\n')

		// Remove newline and parse as integer
		idStr := idInput[:len(idInput)-1]
		if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
			// Check if contact exists and display current contact info
			fmt.Printf("Current contact:\n")
			if err := Store.RenderOne(id); err != nil {
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
			newContact := &contact.Contact{ID: id,
				Name:  newName,
				Email: newEmail}
			Store.Update(newContact)
			break
		}
		fmt.Println("Invalid ID. Please enter a positive number.")
	}
	core(Store)
}

func core(Store contact.Storer) {
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
			addContact(*reader, Store)
		}
	case input == "2\n":
		{
			ListContacts(*reader, Store)
		}
	case input == "3\n":
		{
			handleRemoveContact(*reader, Store)
		}
	case input == "4\n":
		{
			handleUpdateContact(*reader, Store)
		}
	case input == "5\n":
		{
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
