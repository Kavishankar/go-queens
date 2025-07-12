package main

type Row struct {
	cells      []*Cell
	grid       *Grid
	hasQ       bool
	candidates []*Cell
}

func NewRow(g *Grid) *Row {
	return &Row{
		cells:      make([]*Cell, 0, g.dimension),
		grid:       g,
		hasQ:       false,
		candidates: make([]*Cell, 0, g.dimension),
	}
}

func (r *Row) AddCell(cell *Cell) {
	r.cells = append(r.cells, cell)
	r.candidates = append(r.candidates, cell)
}
