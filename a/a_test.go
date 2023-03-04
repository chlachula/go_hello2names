package a

import "testing"

func TestA(t *testing.T) {
    want := "Ann"
    if got := NameA(); got != want {
        t.Errorf("NameA() = %q, want %q", got, want)
    }
}