package myweb

import (
	"testing"
)

func TestsqrtStr(t *testing.T) {
	var want = float64(2)
	if got, _ := sqrtStr("4"); got != want {
		t.Errorf("sqrtStr() = %v want %v", got, want)
	}
}

