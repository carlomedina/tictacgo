package main

import (
	"testing"
)

func TestCheckWinner(t *testing.T) {
	type params struct {
		sum       int
		boardSize int
	}

	var tests = []struct {
		input    params
		expected string
	}{
		{params{3, 1}, ""},
		{params{-3, 3}, "O"},
		{params{4, 4}, "X"},
		{params{-5, 6}, ""},
	}

	for _, test := range tests {
		if output := checkWinner(test.input.sum, test.input.boardSize); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestCheckBoardWinner(t *testing.T) {
	// create games with different board states
	rowX := [][]int{
		{1, 1, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	rowO := [][]int{
		{0, 0, 0},
		{-1, -1, -1},
		{0, 0, 0},
	}
	colX := [][]int{
		{1, 1, 0, 0},
		{1, -1, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}
	colO := [][]int{
		{0, 1, 0, -1},
		{0, -1, 0, -1},
		{0, 1, 0, -1},
		{0, 1, 0, -1},
	}
	diagO := [][]int{
		{-1, 1, 0},
		{0, -1, 0},
		{0, 1, -1},
	}[:]
	diagX := [][]int{
		{0, 1, 1},
		{0, 1, 0},
		{1, 0, 0},
	}
	false1 := [][]int{
		{0, 1, 1},
		{0, 0, 0},
		{1, 0, 0},
	}

	type output struct {
		hasWinner bool
		winner    string
	}
	var tests = []struct {
		input  [][]int
		output output
	}{
		{rowX, output{true, "X"}},
		{rowO, output{true, "O"}},
		{colX, output{true, "X"}},
		{colO, output{true, "O"}},
		{diagX, output{true, "X"}},
		{diagO, output{true, "O"}},
		{false1, output{false, ""}},
	}

	for _, test := range tests {
		if hasWinner, winner := CheckBoard(test.input); hasWinner != test.output.hasWinner && winner != test.output.winner {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.output, output{hasWinner, winner})
		}
	}
}
