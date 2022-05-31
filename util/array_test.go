package util

import "testing"

func TestInArray(t *testing.T) {
	got := InArray("1", []string{"2", "5", "3", "2", "1", "9", "3"})
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
func TestNotInArray(t *testing.T) {
	got := InArray("10", []string{"2", "5", "3", "2", "1", "9", "3"})
	want := false

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
