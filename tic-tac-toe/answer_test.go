package tic_tac_toe

import (
	"strings"
	"testing"
)

/*
x o x
o x o
x o x
 */

func Test_Answer(t *testing.T){
	testCases := []struct{
		name           string
		board          [9]rune
		expectedWinner string
	}{
		{name: "x wins horizontal", board: [9]rune{'x', 'x', 'x', 'o', 'o', 'x', 'o', 'x', 'o'}, expectedWinner: "x"},
		{name: "x wins vertical", board: [9]rune{'x', 'o', 'x', 'x', 'o', 'o', 'x', 'x', 'o'}, expectedWinner: "x"},
		{name: "x wins diagonal", board: [9]rune{'x', 'o', 'x', 'o', 'x', 'o', 'x', 'o', 'x'}, expectedWinner: "x"},

		{name: "o wins horizontal", board: [9]rune{'o', 'o', 'o', 'x', 'x', 'o', 'x', 'o', 'x'}, expectedWinner: "o"},
		{name: "o wins vertical", board: [9]rune{'o', 'x', 'o', 'o', 'x', 'x', 'o', 'o', 'x'}, expectedWinner: "o"},
		{name: "o wins diagonal", board: [9]rune{'o', 'x', 'o', 'x', 'o', 'x', 'o', 'x', 'o'}, expectedWinner: "o"},

	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			w := C(tc.board)
			if !strings.EqualFold(tc.expectedWinner, w){
				t.Errorf("expected winner: %s but was %s\n", tc.expectedWinner, w)
			}
		})
	}


}
