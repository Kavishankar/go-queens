package main

import (
	"bytes"
	"fmt"
)

type Grid struct {
	dimension int
	rows      []*Row
	columns   []*Column
	regions   []*Region
}

func NewGrid(input []byte) (*Grid, error) {
	dim, err := ValidateInput(input)
	if err != nil {
		return nil, err
	}
	g := &Grid{
		dimension: dim,
		rows:      make([]*Row, dim),
		columns:   make([]*Column, dim),
		regions:   make([]*Region, dim),
	}
	var color byte = 'a'
	for idx := range dim {
		g.rows[idx] = NewRow(g)
		g.columns[idx] = NewColumn(g)
		g.regions[idx] = NewRegion(color, g)
		color++
	}
	for idx, val := range input {
		row := g.rows[idx/dim]
		col := g.columns[idx%dim]
		region := g.regions[int(val-'a')]
		c := NewCell(row, col, region, g)
		row.AddCell(c)
		col.AddCell(c)
		region.AddCell(c)
	}
	return g, nil
}

func (g *Grid) Print() {
	fmt.Println(" -------------------------------------")
	for _, row := range g.rows {
		fmt.Print(" | ")
		for colIdx := range g.columns {
			fmt.Printf("%s | ", row.cells[colIdx].String())
		}
		fmt.Println("\n -------------------------------------")
	}
}

func (g *Grid) Solve() bool {
	// TODO: Implement Solve
	return false
}

func (g *Grid) Validate(solvedData []byte) bool {
	var b bytes.Buffer
	for _, row := range g.rows {
		for _, cell := range row.cells {
			b.WriteString(cell.String())
		}
	}
	possibleSol := b.Bytes()
	for idx := range solvedData {
		if solvedData[idx] == 'Q' && possibleSol[idx] != 'Q' {
			return false
		}
		if solvedData[idx] != 'Q' && possibleSol[idx] == 'Q' {
			return false
		}
	}
	return true
}
