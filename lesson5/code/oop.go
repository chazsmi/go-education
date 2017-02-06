// START PACKAGE OMIT
package puzzle

// END PACKAGE OMIT

// START TYPES OMIT
type Piece struct {
	X, Y float64
	Fits bool
}

type Pieces []Piece

// END TYPES OMIT
// START METHODS OMIT
func (pieces Pieces) Draw() Pieces {
	for _, piece := range pieces {
		if piece.Fits {
			pieces = append(pieces, piece)
		}
	}
	return pieces
}

// END METHODS OMIT
func main() {
	// START EG OMIT
	pieces := Pieces{}
	pieces = append(pieces,
		Piece{X: 0.2, Y: 0.5, Fits: true},
		Piece{X: 0.6, Y: 0.1, Fits: true},
		Piece{X: 0.8, Y: 0.2, Fits: false},
	)
	picture := pieces.Draw()
	// END EG OMIT
}
