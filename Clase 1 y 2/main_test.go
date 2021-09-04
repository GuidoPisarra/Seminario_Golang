package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tudai2021.com/model"
)

func TestChangeName(t *testing.T) {

	p := model.NewPerson(1, "alicia")
	changeName(&p, "Agustin")

	assert.Equal(t, p.Name, "Agustin", "los valores no coinciden")
}
