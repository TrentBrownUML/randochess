package rulesetfunctions

import (
	"prushton.com/randochess/v2/board"
)

func CheckersInitBoard(self *board.Board) error {

	for i := range 24 {
		if (i+i/8)%2 == 1 {
			self.Pieces[i].SetPieceType(board.Pawn)
			self.Pieces[i].SetPieceTeam(board.Black)
		}
		if (i+40+i/8)%2 == 0 {
			self.Pieces[i+40].SetPieceType(board.Pawn)
			self.Pieces[i+40].SetPieceTeam(board.White)
		}
	}

	return nil
}

func CheckersGetWinner(self board.Board) board.Team {
	whiteExists := false
	blackExists := false

	for i := range 64 {
		switch self.Pieces[i].GetPieceTeam() {
		case board.White:
			whiteExists = true
		case board.Black:
			blackExists = true
		}
	}

	if whiteExists != blackExists && (whiteExists || blackExists) {
		if whiteExists {
			return board.White
		}
		if blackExists {
			return board.Black
		}
	}
	return board.NoTeam
}

func CheckersMove(self *board.Board, start int, end int, team board.Team) board.Team {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	self.Pieces[end] = self.Pieces[start]
	self.Pieces[start].SetPieceTeam(board.NoTeam)

	// handle turning into kings
	white_piece_at_last_rank := end/self.Height == 0 && self.Pieces[end].GetPieceTeam() == board.White
	black_piece_at_first_rank := end/self.Height == 7 && self.Pieces[end].GetPieceTeam() == board.Black
	if white_piece_at_last_rank || black_piece_at_first_rank {
		self.Pieces[end].SetPieceType(board.King)
	}

	// handle jumping over other piece
	if Abs(delta_x) == 2 && Abs(delta_y) == 2 {
		hopped_piece_x := -delta_x / 2
		hopped_piece_y := -(delta_y / 2) * 8
		// fmt.Printf("%d, %d | %d %d", hopped_piece_x, hopped_piece_y, start, start+hopped_piece_x+hopped_piece_y)
		hopped_piece := self.Pieces[hopped_piece_x+hopped_piece_y+start]

		if hopped_piece.GetPieceTeam() != board.NoTeam {
			self.Pieces[hopped_piece_x+hopped_piece_y+start].SetPieceTeam(board.NoTeam)
			return team
		}
	}

	return team.OtherTeam()
}

func CheckersPawn(self board.Board, start int, end int) bool {
	var delta_y int = start/self.Height - end/self.Height

	if delta_y < 0 && self.Pieces[start].GetPieceTeam() == board.White {
		return false
	}

	if delta_y > 0 && self.Pieces[start].GetPieceTeam() == board.Black {
		return false
	}

	return CheckersKing(self, start, end)
}

func CheckersKing(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height
	var starting_piece board.Piece = self.Pieces[start]

	if Abs(delta_x) != Abs(delta_y) {
		return false
	}

	if Abs(delta_y) == 1 {
		return true
	}

	if Abs(delta_y) == 2 { // we are hopping over a piece
		// negative for some reason
		hopped_piece_x := -delta_x / 2
		hopped_piece_y := -(delta_y / 2) * 8
		// fmt.Printf("%d, %d | %d %d", hopped_piece_x, hopped_piece_y, start, start+hopped_piece_x+hopped_piece_y)
		hopped_piece := self.Pieces[hopped_piece_x+hopped_piece_y+start]

		return hopped_piece.GetPieceTeam() != board.NoTeam && hopped_piece.GetPieceTeam() != starting_piece.GetPieceTeam()
		// return false
	}
	return false
}
