package contact

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Storer interface {
	Save(c *Contact) error
	Render() error
	RenderOne(id int) error
	Delete(id int) error
	Update(c *Contact) error
}

type MemoryStore struct {
	Contacts map[int]*Contact
}
