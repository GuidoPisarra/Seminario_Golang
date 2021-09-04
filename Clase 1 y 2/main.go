package main

import (
	"fmt"

	"tudai2021.com/model"
)

func changeName(p *model.Person, name string) {
	p.Name = name
}
func main() {
	p := model.NewPerson(0, "Guido")
	fmt.Println(p)
	changeName(&p, "Pepe")
	fmt.Println(p)
}
