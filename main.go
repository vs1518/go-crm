package main

import (
	"minicrm/cmd"
	"minicrm/contact"
)

var (
	nextid = 0
)

func main() {
	Store := &contact.MemoryStore{
		Contacts: make(map[int]*contact.Contact),
	}
	core(Store)
	cmd.Execute()

}
