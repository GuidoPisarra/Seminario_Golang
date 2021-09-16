package model

import (
	"fmt"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

type Caracter interface {
	cumple(r Result) (*Result, error)
}

func (c *Result) CrearResult(cadena string) (*Result, error) {
	t, err := compruebaTipo(cadena[0:2], cadena[4:])
	if err != nil {
		return &Result{"", "", 0}, fmt.Errorf("this is an %s error", "no valid")
	}
	l, err := strconv.Atoi(cadena[2:4])
	if err != nil {
		return &Result{"", "", 0}, fmt.Errorf("this is an %s error", "no valid")
	}
	v := cadena[4:]
	l, err = lon(l, v)
	if err != nil {
		return &Result{"", "", 0}, fmt.Errorf("this is an %s error", "no valid")
	}
	r := NewResult(t, v, l)
	return &r, nil
}

func lon(l int, v string) (int, error) {
	if l == len(v) {
		return l, nil
	}
	return 0, fmt.Errorf("this is an %s error", "no valid")
}

func compruebaTipo(t string, c string) (string, error) {
	if t == "NN" {
		l, err := strconv.Atoi(c)
		if err == nil && l > 0 {
			return t, err
		}
		t = ""
		return t, err
	}
	if t == "TX" {
		return t, nil
	}

	return t, fmt.Errorf("this is an %s error", "no valid")

}

func NewResult(typ string, value string, length int) Result {
	return Result{typ, value, length}
}
