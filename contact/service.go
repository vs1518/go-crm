package contact

import "fmt"

func (store MemoryStore) Save(contact *Contact) error {
	store.Contacts[contact.ID] = contact
	return nil
}

func (store MemoryStore) Render() error {
	fmt.Println("+------+-----------+----------------+")
	fmt.Println("| ID   | Name      | Email          |")
	fmt.Println("+------+-----------+----------------+")
	for _, contact := range store.Contacts {
		fmt.Printf("| %-4d | %-10s | %-14s |\n", contact.ID, contact.Name, contact.Email)
	}
	fmt.Println("+------+-----------+----------------+")
	return nil
}

func (store MemoryStore) RenderOne(id int) error {
	contact, exists := store.Contacts[id]
	if !exists {
		return fmt.Errorf("contact not found")
	}
	fmt.Println("+------+-----------+----------------+")
	fmt.Println("| ID   | Name      | Email          |")
	fmt.Println("+------+-----------+----------------+")
	fmt.Printf("| %-4d | %-10s | %-14s |\n", contact.ID, contact.Name, contact.Email)
	fmt.Println("+------+-----------+----------------+")
	return nil
}
func (store MemoryStore) Delete(id int) error {
	_, exists := store.Contacts[id]
	if !exists {
		return fmt.Errorf("contact not found with ID: %d", id)
	}
	delete(store.Contacts, id)
	return nil
}

func (store MemoryStore) Update(c *Contact) error {
	_, exists := store.Contacts[c.ID]
	if !exists {
		return fmt.Errorf("contact not found")
	}
	store.Contacts[c.ID] = c
	return nil
}
