package rulesetfunctions

import "prushton.com/randochess/v2/board"

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

	if whiteExists != blackExists && whiteExists || blackExists {
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
	return team.OtherTeam()
}

func CheckersPawn(self board.Board, start int, end int) bool
