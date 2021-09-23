package main

import "testing"

func TestHelloWorld(t *testing.T) {
	msg := HelloWorld()
	if msg != "Hello World" {
		t.Fatalf(`HelloWorld() = %q, want match for %#q, nil`, msg, "Hello World")
	}
}
