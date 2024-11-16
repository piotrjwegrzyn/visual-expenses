package main

import "testing"

func Test_helloWorld(t *testing.T) {
	expected := "Hello, World!"

	actual := helloWorld()

	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}
