package b

import "testing"

func TestB(t *testing.T) {
    want := "Bob"
    if got := NameB(); got != want {
        t.Errorf("NameB() = %q, want %q", got, want)
    }
}