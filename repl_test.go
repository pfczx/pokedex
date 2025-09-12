package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "hello world",
            expected: []string{"hello", "world"},
        },
        {
            input:    "hei tei",
            expected: []string{"hei", "tei"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Length mismatch: expected %v, got %v", len(c.expected), len(actual))
            continue
        }
        
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("Word mismatch at index %d: expected %s, got %s", i, c.expected[i], actual[i])
            }
        }
    }
}
