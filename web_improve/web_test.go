package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqrtStr(t *testing.T) {

	tests := [] struct {
		name string
		input string
		output float64
		errmsg string
	}{
		{
			name :"testing valid digit",
			input : "4",
			output: 2 ,
			errmsg: "",
		},
		{
			name: "testing letter as an input - func should return an error",
			input: "g",
			output: nil ,
			errmsg: "strconv.ParseFloat: parsing \"g\": invalid syntax",
		},
		{
			name: "testing case with no input symbol",
			input: "",
			output: nil,
			errmsg: "strconv.ParseFloat: parsing \"\": invalid syntax",
		},

	}

	for _, tt := range tests{
			res, err := sqrtStr(tt.input)
			errms := err.Error()
			assert.Equal(t, tt.errmsg, errms, "sqrtStr returns unexpected error")
			assert.Equal(t, tt.output, res, "sqrtStr returns unexpected value")



			}
}




