package main

type Column struct {
	cells      []*Cell
	grid       *Grid
	hasQ       bool
	candidates []*Cell
}

func NewColumn(g *Grid) *Column {
	return &Column{
		cells:      make([]*Cell, 0, g.dimension),
		grid:       g,
		hasQ:       false,
		candidates: make([]*Cell, 0, g.dimension),
	}
}

func (c *Column) AddCell(cell *Cell) {
	c.cells = append(c.cells, cell)
	c.candidates = append(c.candidates, cell)
}
