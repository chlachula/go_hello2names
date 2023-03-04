package main

import "testing"

func Test(t *testing.T) {
    want := "Hello Ann & Bob. Good to see you."
    if got := text(); got != want {
        t.Errorf("nameA() = %q, want %q", got, want)
    }
}