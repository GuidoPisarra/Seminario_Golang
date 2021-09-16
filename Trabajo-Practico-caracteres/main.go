package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	c := "TX02AA"
	var res model.Result
	var err error
	var r *model.Result
	r, err = res.CrearResult(c)
	if err == nil {
		fmt.Println(r)
	}

}

func cumple(v interface{}) (bool, error) {
	return true, nil
}
