package board

import (
	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	panic("not implemented") 
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	pos := c[x][y]
	if val != nil {
		return val
	}
	return nil
	
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	// TODO: Implement
	if c.PieceAt(from) != nil {
		x, _ := from.Coord(0)
		y, _ := from.Coord(1)
		pos := c[x][y]
		return c.PlacePieceAt(pos, to)
	}
	else{
		return fmt.Errorf("No piece at %v", from.String())
	}


}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	if c.PieceAt(at) != nil {
		return fmt.Errorf("Piece at %v", at.String())
	}	
	else{
		x, _ := at.Coord(0)
		y, _ := at.Coord(1)
		c[x][y] = pos
		return nil
	}
}
