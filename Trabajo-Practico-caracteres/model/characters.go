package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

func (c *Result) CrearResult(cadena string) (*Result, error) {
	t, err1 := compruebaTipo(cadena[0:2], cadena[4:])
	l, err2 := strconv.Atoi(cadena[2:4])
	v := cadena[4:]
	k, err3 := lon(l, v)
	if err1 == nil && err2 == nil && err3 == nil {
		r := NewResult(t, v, k)
		return &r, nil
	}
	return &Result{"", "", 0}, errors.New("no valid")
}

func lon(l int, v string) (int, error) {
	if l == len(v) {
		return l, nil
	}
	return 0, errors.New("no valid")
}

func compruebaTipo(t string, c string) (string, error) {
	if t == "NN" {
		_, err := strconv.Atoi(c)
		if err == nil {
			return t, err
		}
		t = ""
		return t, err
	}
	if t == "TX" {
		return t, nil
	}

	return t, errors.New("no valid")

}

func NewResult(typ string, value string, length int) Result {
	return Result{typ, value, length}
}
