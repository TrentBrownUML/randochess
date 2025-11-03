package rulesetfunctions

import "prushton.com/randochess/v2/board"

func CheckersInitBoard(self *board.Board) error {

	for i := range 24 {
		if i%2 == 1 {
			self.Pieces[i].SetPieceType(board.Pawn)
			self.Pieces[i].SetPieceTeam(board.Black)
		}
		if (i+40)%2 == 0 {
			self.Pieces[i].SetPieceType(board.Pawn)
			self.Pieces[i].SetPieceTeam(board.White)
		}
	}

	return nil
}
