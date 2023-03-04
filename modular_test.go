package go_hello2names

import "testing"

func TestA(t *testing.T) {
	want := "Good to see you."
	if got := Message(); got != want {
		t.Errorf("Message() = %q, want %q", got, want)
	}
}
