package main

type Region struct {
	color      byte
	cells      []*Cell
	grid       *Grid
	hasQ       bool
	candidates []*Cell
}

func NewRegion(color byte, g *Grid) *Region {
	return &Region{
		color:      color,
		cells:      make([]*Cell, 0, g.dimension),
		grid:       g,
		hasQ:       false,
		candidates: make([]*Cell, 0, g.dimension),
	}
}

func (r *Region) AddCell(cell *Cell) {
	r.cells = append(r.cells, cell)
	r.candidates = append(r.candidates, cell)
}
