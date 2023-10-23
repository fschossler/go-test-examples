package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello World!"

	if got != want {
		t.Errorf("The printing %q is wrong, must return %q", got, want)
	}
}
