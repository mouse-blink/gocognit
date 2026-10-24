package main

import (
	"testing"

	"github.com/uudashr/gocognit"
)

func TestAverage(t *testing.T) {
	stats := []gocognit.Stat{
		{FuncName: "A", Complexity: 2},
		{FuncName: "B", Complexity: 4},
		{FuncName: "C", Complexity: 9},
	}

	if got, want := average(stats), 5.0; got != want {
		t.Errorf("average = %v, want %v", got, want)
	}
}
