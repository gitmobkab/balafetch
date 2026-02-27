package strings_helpers

import (
	"testing"
)

func TestTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"hElLo WoRLD", "Hello World"},
		{"goLang", "Golang"},
	}
	for _, test := range tests {
		result := TitleCase(test.input)
		if result != test.expected {
			t.Errorf("TitleCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello world"},
		{"HELLO WORLD", "hello world"},
		{"hElLo WoRLD", "hello world"},
		{"goLang", "golang"},
	}
	for _, test := range tests {
		result := LowerCase(test.input)
		if result != test.expected {
			t.Errorf("LowerCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	strs := []string{"Hello", "World", "GoLang"}
	tests := []struct {
		str      string
		expected bool
	}{
		{"hello", true},
		{"WORLD", true},
		{"golang", true},
		{"Python", false},
	}

	for _, test := range tests {
		result := ContainsIgnoreCase(strs, test.str)
		if result != test.expected {
			t.Errorf("ContainsIgnoreCase(%v, %q) = %v; want %v", strs, test.str, result, test.expected)
		}
	}
}