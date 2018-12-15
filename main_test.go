package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello World!" {
		t.Errorf("got %s expected %s", HelloWorld(), "Hello World!")
	}
}
