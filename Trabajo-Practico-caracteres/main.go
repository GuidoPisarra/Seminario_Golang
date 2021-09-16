package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	var cad string

	fmt.Print("Ingrese una cadena ")
	fmt.Scan(&cad)
	var res model.Result
	var err error
	var r *model.Result
	r, err = res.CrearResult(cad)
	if err == nil {
		fmt.Println(r)
	} else {
		fmt.Println(err)
	}

}
