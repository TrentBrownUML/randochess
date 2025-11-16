package rulesetfunctions

import (
	"fmt"

	"prushton.com/randochess/v2/board"
)

func DefaultMove(self *board.Board, start int, end int, team board.Team) board.Team {
	self.Pieces[end] = self.Pieces[start]
	self.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Pieces[end].SetPieceMoved()
	return team.OtherTeam()
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

func DefaultRook(self board.Board, start int, end int) ([]int, []int) {
	var validMoveLocations []int = make([]int, 0)
	var directions [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// construct an array of spaces where the piece can move.
	// Iterate over every direction and look until we reach the edge of the board or a piece
	for _, direction := range directions {
		var distance int = 1
		var reachedLimit bool = false

		for !reachedLimit {

			destination := start + direction[0]*distance + direction[1]*self.Width*distance

			if CheckLineOfSight(self, start, destination) {
				validMoveLocations = append(validMoveLocations, destination)
			} else {
				reachedLimit = true
			}

			distance += 1
		}
	}

	// most pieces can take at the same spots they can move to, so i just return them both
	return validMoveLocations, validMoveLocations
}

func DefaultKnight(self board.Board, start int, end int) ([]int, []int) {
	var validMoveLocations []int = make([]int)
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
