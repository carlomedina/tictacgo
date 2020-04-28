package main

import (
	"testing"
)

func TestValidatePiece(t *testing.T) {
	type output struct {
		value   interface{}
		isValid bool
		res     string
	}

	var tests = []struct {
		input    string
		expected output
	}{
		{"X", output{"X", true, ""}},
		{"O", output{"O", true, ""}},
		{"", output{"", false, ""}},
		{"1", output{"", false, ""}},
		{"a", output{"", false, ""}},
	}

	for _, test := range tests {
		if value, ok, _ := validatePiece(test.input); value.(string) != test.expected.value || ok != test.expected.isValid {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, []interface{}{value.(string), ok})
		}
	}
}

func TestValidateBoardSize(t *testing.T) {
	type output struct {
		value   interface{}
		isValid bool
		res     string
	}

	var tests = []struct {
		input    string
		expected output
	}{
		{"3", output{3, true, ""}},
		{"-1", output{0, false, ""}},
		{"100", output{0, false, ""}},
		{"5", output{5, true, ""}},
		{"a", output{0, false, ""}},
	}

	for _, test := range tests {
		if value, ok, _ := validateBoardSize(test.input); value.(int) != test.expected.value || ok != test.expected.isValid {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, []interface{}{value.(int), ok})
		}
	}
}

func TestValidateNumPlayers(t *testing.T) {
	type output struct {
		value   interface{}
		isValid bool
		res     string
	}

	var tests = []struct {
		input    string
		expected output
	}{
		{"3", output{0, false, ""}},
		{"1", output{1, true, ""}},
		{"2", output{2, true, ""}},
		//{"1,", output{0, false, ""}}, // TODO: need to handle this
		{"a", output{0, false, ""}},
	}

	for _, test := range tests {
		if value, ok, _ := validateNumPlayers(test.input); value.(int) != test.expected.value || ok != test.expected.isValid {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, []interface{}{value.(int), ok})
		}
	}
}

func TestValidateEndGameResponse(t *testing.T) {
	type output struct {
		value   interface{}
		isValid bool
		res     string
	}

	var tests = []struct {
		input    string
		expected output
	}{
		{"restart", output{"restart", true, ""}},
		{"end", output{"end", true, ""}},
		{"new", output{"new", true, ""}},
		{"1", output{"", false, ""}},
		{"a", output{"", false, ""}},
		{"$%#$4", output{"", false, ""}},
	}

	for _, test := range tests {
		if value, ok, _ := validateEndGameResponse(test.input); value.(string) != test.expected.value || ok != test.expected.isValid {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, []interface{}{value.(string), ok})
		}
	}
}
