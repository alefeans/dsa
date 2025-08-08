package main

import "testing"

const errorMsg = "sent value %v, expected value %v, found %v"

func TestCheckWeird(t *testing.T) {
	t.Parallel()
	tests := []struct {
		sent     int32
		expected string
	}{
		{3, "Weird"},
		{8, "Weird"},
		{20, "Weird"},
		{24, "Not Weird"},
		{4, "Not Weird"},
	}

	for _, test := range tests {
		t.Logf("Data test: %v", test)
		value := checkWeird(test.sent)
		if value != test.expected {
			t.Errorf(errorMsg, test.sent, test.expected, value)
		}
	}
}
