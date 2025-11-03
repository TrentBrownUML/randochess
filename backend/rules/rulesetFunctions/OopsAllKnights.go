package rulesetfunctions

import (
	"fmt"

	"prushton.com/randochess/v2/board"
)

func OopsAllKnightsInitBoard(self *board.Board) error {
	if self.Height%2 == 1 || self.Width%2 == 1 {
		return fmt.Errorf("Cannot init board with odd width or height")
	}
	heightOffset := (self.Height - 8) / 2
	widthOffset := (self.Width - 8) / 2

	offset := heightOffset*self.Width + widthOffset

	backRow := [8]board.PieceType{board.Knight, board.Knight, board.Knight, board.Knight, board.King, board.Knight, board.Knight, board.Knight}

	for i := range 8 {
		self.Pieces[offset+i].SetPieceTeam(board.Black)
		self.Pieces[offset+i].SetPieceType(backRow[i])

		self.Pieces[offset+i+self.Width].SetPieceTeam(board.Black)
		self.Pieces[offset+i+self.Width].SetPieceType(board.Knight)

		self.Pieces[offset+i+self.Width*6].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*6].SetPieceType(board.Knight)

		self.Pieces[offset+i+self.Width*7].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*7].SetPieceType(backRow[i])
	}

	return nil
}
