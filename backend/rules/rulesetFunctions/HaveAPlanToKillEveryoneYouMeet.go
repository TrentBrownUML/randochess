package rulesetfunctions

import "prushton.com/randochess/v2/board"

func PlanBishop(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	// cant move well, but can take without LOS
	if self.Pieces[end].GetPieceTeam() == board.NoTeam {
		return Abs(delta_x) <= 1 && Abs(delta_y) <= 1
	} else {
		return Abs(delta_x) == Abs(delta_y)
	}
}
