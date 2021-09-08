package main

import (
	"fmt"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

type caracter interface {
	cumple(r Result)
}

func main() {
	c := "NN0431A3"
	var r Result
	r = crearResul(c)
	fmt.Println(r)

}

func crearResul(cadena string) Result {
	var o Result
	o.Type = cadena[0:2]

	o.Length, _ = strconv.Atoi(cadena[2:4])
	o.Value = cadena[4:]
	cumple(o)

	return o
}

func compruebaTipo(t string) bool {
	return (t == "NN" || t == "TX")
}

func cumple(v interface{}) bool {
	s := v.(Result)
	d := s.Value

	if s.Type == "NN" {
		gr, ok := strconv.Atoi(d)
		if ok == nil {
			fmt.Println(s)
		} else {
			fmt.Println("no es numero", gr)
		}
	}

	return true /* r.Length-4 == len(r.Value) */
}
