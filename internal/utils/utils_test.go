package utils

import (
	"testing"
)

func TestSome(t *testing.T) {
	cases := []struct {
		msg      string
		keys     []string
		expected bool
	}{
		{"Fix bug in feature X", []string{"BUG-123", "FEATURE-456"}, false},
		{"Implement FEATURE-456", []string{"BUG-123", "FEATURE-456"}, true},
		{"Refactor code", []string{"BUG-123", "FEATURE-456"}, false},
	}

	for _, c := range cases {
		got := Some(&c.msg, &c.keys)
		if got != c.expected {
			t.Errorf("Some(%q, %v) received \"%v\", expected \"%v\"", c.msg, c.keys, got, c.expected)
		}
	}
}

func TestGetCommitTitle(t *testing.T) {
	cases := []struct {
		msg      string
		expected string
	}{
		{"Fix bug in feature X\nMore details about the commit", "Fix bug in feature X"},
		{"Implement FEATURE-456\nMore details about the commit\nEven more details", "Implement FEATURE-456"},
		{"Refactor code", "Refactor code"},
	}

	for _, c := range cases {
		got := GetCommitTitle(&c.msg)
		if got != c.expected {
			t.Errorf("GetCommitTitle(%q) received \"%v\", expected \"%v\"", c.msg, got, c.expected)
		}
	}
}
