package leetcode

import "testing"

func Test_uniqUser(t *testing.T) {
	users := [][]string{
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Alice", "alice@example.com"},
		{"Charlie", "alice@example.com"},
		{"David", "david@example.com"},
		{"Eve", "eve@example.com"},
		{"Bob", "bob@example.com"},
	}

	want := 4
	got := uniqUser(users)

	if got != want {
		t.Errorf("uniqUser() = %d, want %d", got, want)
	}
}

func Test_uniqUser_EmptyInput(t *testing.T) {
	users := [][]string{}

	want := 0
	got := uniqUser(users)

	if got != want {
		t.Errorf("uniqUser() = %d, want %d", got, want)
	}
}

func Test_uniqUser_SingleUser(t *testing.T) {
	users := [][]string{
		{"Alice", "alice@example.com"},
	}

	want := 1
	got := uniqUser(users)

	if got != want {
		t.Errorf("uniqUser() = %d, want %d", got, want)
	}
}

func Test_uniqUser_NoDuplicates(t *testing.T) {
	users := [][]string{
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
		{"David", "david@example.com"},
		{"Eve", "eve@example.com"},
	}

	want := 5
	got := uniqUser(users)

	if got != want {
		t.Errorf("uniqUser() = %d, want %d", got, want)
	}
}
