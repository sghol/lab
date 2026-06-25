package main

import (
	"fmt"
	"testing"
)

func Test_Hello(t *testing.T) {
	fmt.Println("First test")
	expected := "hello"
	result := Hello()
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)

	}

}
