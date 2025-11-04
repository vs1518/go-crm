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
	load(id int) (*contact, error)
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

func (store MemoryStore) load(id int) (*contact, error) {
	contact, exists := store.contacts[id]
	if !exists {
		return nil, fmt.Errorf("contact not found")
	}
	return contact, nil
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
