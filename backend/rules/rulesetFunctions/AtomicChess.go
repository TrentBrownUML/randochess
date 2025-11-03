package rulesetfunctions

import "prushton.com/randochess/v2/board"

func AtomicChessMove(self *board.Board, start int, end int) {
	if self.Pieces[end].GetPieceTeam() == board.NoTeam {
		self.Pieces[end] = self.Pieces[start]
		self.Pieces[start].SetPieceTeam(board.NoTeam)
		self.Pieces[end].SetPieceMoved()
	} else {
		self.Pieces[end].SetPieceTeam(board.NoTeam)
		self.Pieces[start].SetPieceTeam(board.NoTeam)
	}
}
