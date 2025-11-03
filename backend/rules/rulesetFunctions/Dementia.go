package rulesetfunctions

import (
	"math/rand"

	"prushton.com/randochess/v2/board"
)

func DementiaMove(self *board.Board, start int, end int, team board.Team) board.Team {
	self.Pieces[end] = self.Pieces[start]
	self.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Pieces[end].SetPieceMoved()

	r := rand.Intn(96)
	if r < 64 && self.Pieces[r].GetPieceType() != board.King {
		self.Pieces[r].SetPieceTeam(board.NoTeam)
	}

	return team.OtherTeam()
}
