package model

type Person struct {
	ID   int
	Name string
}

func NewPerson(id int, name string) Person {
	return Person{id, name}
}
