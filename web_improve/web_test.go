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
		wantErr bool
	 	errmsg string
	}{
		{
			name :"testing valid digit",
			input : "4",
			output: 2 ,
			wantErr: false,
		},
		{
			name: "testing letter as an input - func should return an error",
			input: "g",
			output: 0 ,
			wantErr: true,
			errmsg: "strconv.ParseFloat: parsing \"g\": invalid syntax",
		},
		{
			name: "testing case with no input symbol",
			input: "",
			output: 0 ,
			wantErr: true,
			errmsg: "strconv.ParseFloat: parsing \"\": invalid syntax",
		},

	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := sqrtStr(tt.input)

			assert.Equal(t, tt.output, res, "sqrtStr returns unexpected value")
			if tt.wantErr {
				assert.Equal(t, tt.errmsg, err.Error(), "sqrtStr returns unexpected error")
			}else {
				assert.Equal(t, nil, err, "sqrtStr returns unexpected error")
			}
		})
	}
}
