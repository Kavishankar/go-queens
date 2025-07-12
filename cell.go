package main

type Cell struct {
	row    *Row
	column *Column
	region *Region
	grid   *Grid
	isQ    bool
	isX    bool
}

func NewCell(r *Row, c *Column, reg *Region, g *Grid) *Cell {
	return &Cell{
		row:    r,
		column: c,
		region: reg,
		grid:   g,
		isQ:    false,
		isX:    false,
	}
}

func (c *Cell) String() string {
	switch {
	case c.isQ:
		return "Q"
	case c.isX:
		return "X"
	default:
		return string(c.region.color)
	}
}
