package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearResul(t *testing.T) {

	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"AA04000A", false, "", "", 0},
		/*{"NN04000A", false, "", "", 0},
		{"NN0A000A", false, "", "", 0}, */
	}

	for _, testData := range cases {
		var res Result
		r, err := res.CrearResult(testData.Input)

		assert.Equal(t, err == nil, testData.Success)
		if r != nil {
			assert.Equal(t, testData.Type, r.Type)
			assert.Equal(t, testData.Value, r.Value)
			assert.Equal(t, testData.Length, r.Length)

		}
	}

}
