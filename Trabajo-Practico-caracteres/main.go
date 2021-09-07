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

func main() {
	c := "AA0431234"
	var r Result
	r.Length = obtieneCantidad(c, 3, 4)
	r.Value = obtenerCadena(c, r.Length)
	r.Type = c[0:2]
	fmt.Println(r.Length)
	fmt.Println(r.Value)
	fmt.Println(r.Type)

}

func obtieneCantidad(cadena string, inicio int, fin int) int {
	var cant string
	var c int
	for i := 0; i < len(cadena); i++ {
		if (i >= inicio-1) && (i <= fin-1) {
			cant += string(cadena[i])

		}
	}
	c, _ = strconv.Atoi(cant)
	return c
}

func obtenerCadena(cadena string, cant int) string {
	var valor string
	for i := 0; i < cant; i++ {
		valor += string(cadena[i+5])
	}
	return valor
}
