package main

import (
	"fmt"
)
type contact struct {
	ID    int
	name  string
	email string
}

type storer interface {
	save(c *contact) error
	render() error
	renderOne(id int) error
	delete(id int) error
	update(c *contact) error
}

type MemoryStore struct {
	contacts map[int]*contact
}

func (store MemoryStore) save(contact *contact) error {
	store.contacts[contact.ID] = contact
	return nil
}

func (store MemoryStore) render() error {
	fmt.Println("+------+-----------+----------------+")
	fmt.Println("| ID   | Name      | Email          |")
	fmt.Println("+------+-----------+----------------+")
	for _, contact := range store.contacts {
		fmt.Printf("| %-4d | %-10s | %-14s |\n", contact.ID, contact.name, contact.email)
	}
	fmt.Println("+------+-----------+----------------+")
	return nil
}

func (store MemoryStore) renderOne(id int) error {
	contact, exists := store.contacts[id]
	if !exists {
		return fmt.Errorf("contact not found")
	}
	fmt.Println("+------+-----------+----------------+")
	fmt.Println("| ID   | Name      | Email          |")
	fmt.Println("+------+-----------+----------------+")
	fmt.Printf("| %-4d | %-10s | %-14s |\n", contact.ID, contact.name, contact.email)
	fmt.Println("+------+-----------+----------------+")
	return nil
}
func (store MemoryStore) delete(id int) error {
	_, exists := store.contacts[id]
	if !exists {
		return fmt.Errorf("contact not found")
	}
	delete(store.contacts, id)
	return nil
}

func (store MemoryStore) update(c *contact) error {
	_, exists := store.contacts[c.ID]
	if !exists {
		return fmt.Errorf("contact not found")
	}
	store.contacts[c.ID] = c
	return nil
}
