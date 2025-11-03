package rulesetfunctions

import "prushton.com/randochess/v2/board"

func Knook(self board.Board, start int, end int) bool {
	return DefaultKnight(self, start, end) || DefaultRook(self, start, end)
}
