package board

type Board struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Pieces []Piece `json:"pieces"`
}

func New(w int, h int) Board {
	return Board{
		Width:  int(w),
		Height: int(h),
		Pieces: make([]Piece, w*h),
	}
}
