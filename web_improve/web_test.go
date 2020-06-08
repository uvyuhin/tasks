package main

import (
	"reflect"
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
			errmsg: nil,
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
			got := struct {float64; string }{res, errms}
			want := struct {float64; string}{tt.output, tt.errmsg}
			if !reflect.DeepEqual(got, want){
				t.Fatalf("test %d: expected: %v, got: %v",tt.name, want, got)
			}

			}

	}



