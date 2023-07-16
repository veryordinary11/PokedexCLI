package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(actual),
				len(c.expected),
			)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("The values are not equal: %v vs %v",
					actual[i],
					c.expected[i],
				)
			}
		}
	}
}
