package main

import (
	"testing"
)

func TestSqrtStr(t *testing.T) {

	tests := [] struct {
		name string
		input string
		output float64
		errmsg error
	}{
		{
			name :"testing valid digit",
			input : "4",
			output: 2 ,
			errmsg: nil ,
		},
		{
			name: "testing letter as an input - func should return an error",
			input: "g",
			output: nil ,
			errmsg: ""
		},

	}

	//testing valid digit
	var want float64 = 2
	if got, _ := sqrtStr("4"); got != want {
		t.Errorf("sqrtStr() = %v want %v", got, want)
	}

	//testing letter - func should return an error
	if _ , got1 := sqrtStr("d"); got1 == nil {
		t.Errorf("sqrtStr() = %v want %v", got1, nil)
	}

	//testing case with no input symbol
	if _ , got2 := sqrtStr(""); got2 == nil {
		t.Errorf("sqrtStr() = %v want %v", got2, nil)
	}
}

