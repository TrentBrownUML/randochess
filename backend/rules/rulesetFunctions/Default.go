package rulesetfunctions

import (
	"fmt"

	"prushton.com/randochess/v2/board"
)

func DefaultMove(self *board.Board, start int, end int) {
	self.Pieces[end] = self.Pieces[start]
	self.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Pieces[end].SetPieceMoved()
}

func DefaultGetWinner(self board.Board) board.Team {
	WhiteInPlay := false
	BlackInPlay := false

	for i := range self.Width * self.Height {
		if self.Pieces[i].GetPieceType() == board.King {
			if self.Pieces[i].GetPieceTeam() == board.White {
				WhiteInPlay = true
			}
			if self.Pieces[i].GetPieceTeam() == board.Black {
				BlackInPlay = true
			}
		}
	}

	if WhiteInPlay != BlackInPlay && (WhiteInPlay || BlackInPlay) {
		if WhiteInPlay {
			return board.White
		}
		return board.Black
	}
	return board.NoTeam
}

func DefaultInitBoard(self *board.Board) error {
	if self.Height%2 == 1 || self.Width%2 == 1 {
		return fmt.Errorf("Cannot init board with odd width or height")
	}
	heightOffset := (self.Height - 8) / 2
	widthOffset := (self.Width - 8) / 2

	offset := heightOffset*self.Width + widthOffset

	backRow := [8]board.PieceType{board.Rook, board.Knight, board.Bishop, board.Queen, board.King, board.Bishop, board.Knight, board.Rook}

	for i := range 8 {
		self.Pieces[offset+i].SetPieceTeam(board.Black)
		self.Pieces[offset+i].SetPieceType(backRow[i])

		self.Pieces[offset+i+self.Width].SetPieceTeam(board.Black)
		self.Pieces[offset+i+self.Width].SetPieceType(board.Pawn)

		self.Pieces[offset+i+self.Width*6].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*6].SetPieceType(board.Pawn)

		self.Pieces[offset+i+self.Width*7].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*7].SetPieceType(backRow[i])
	}

	return nil
}

func DefaultPawn(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	// moving backwards
	if delta_y < 0 && self.Pieces[start].GetPieceTeam() == board.White {
		return false
	}

	if delta_y > 0 && self.Pieces[start].GetPieceTeam() == board.Black {
		return false
	}

	// moving 1 space or 2 on first turn
	if (Abs(delta_y) == 1 && delta_x == 0) || (Abs(delta_y) == 2 && !self.Pieces[start].GetPieceMoved()) {
		return self.Pieces[end].GetPieceTeam() == board.NoTeam && CheckLineOfSight(self, start, end)
	}

	// taking
	if Abs(delta_y) == 1 && Abs(delta_x) == 1 && self.Pieces[end].GetPieceTeam() != board.NoTeam {
		return true
	}

	return false
}

func DefaultRook(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return ((delta_x == 0) != (delta_y == 0)) && CheckLineOfSight(self, start, end)
}

func DefaultKnight(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return (Abs(delta_x) == 1 && Abs(delta_y) == 2) || (Abs(delta_x) == 2 && Abs(delta_y) == 1)
}
func DefaultBishop(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return Abs(delta_x) == Abs(delta_y) && CheckLineOfSight(self, start, end)
}

func DefaultKing(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return delta_x >= -1 && delta_x <= 1 && delta_y >= -1 && delta_y <= 1
}

func DefaultQueen(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	if delta_x == 0 || delta_y == 0 {
		return CheckLineOfSight(self, start, end)
	}

	if Abs(delta_x) == Abs(delta_y) {
		return CheckLineOfSight(self, start, end)
	}

	return false
}
