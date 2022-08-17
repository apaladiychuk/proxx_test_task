package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Location struct {
	Row    int
	Column int
}

type Cell struct {
	blackHole               bool
	amountAdjacentBlackHole int
	isOpened                bool
}

type Board struct {
	Rows       int
	Columns    int
	BlackHoles int
	Map        map[Location]*Cell
}

func NewBoard(rows, columns, blackHoles int) (*Board, error) {
	if blackHoles > rows*columns {
		return nil, fmt.Errorf("no math blackholes")
	}
	b := Board{
		Rows:       rows,
		Columns:    columns,
		BlackHoles: blackHoles,
		Map:        make(map[Location]*Cell),
	}
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < blackHoles; i++ {
		loc := Location{
			Row:    int(rnd.Float64() * float64(rows)),
			Column: int(rnd.Float64() * float64(rows)),
		}
		if _, ok := b.Map[loc]; ok {
			i--
			continue
		}

		b.Map[loc] = &Cell{
			blackHole:               true,
			amountAdjacentBlackHole: 0,
		}
		b.fillAdjacent(&loc)
	}
	return &b, nil
}
func (b *Board) adjacentIndexes(i, max int) (start, finish int) {
	start = i - 1
	if start < 0 {
		start = 0
	}
	finish = i + 1
	if finish > max-1 {
		finish = max - 1
	}
	return
}

func (b *Board) fillAdjacent(loc *Location) {
	startColumn, finishColumn := b.adjacentIndexes(loc.Column, b.Columns)
	startRow, finishRow := b.adjacentIndexes(loc.Row, b.Rows)

	for r := startRow; r <= finishRow; r++ {
		for c := startColumn; c <= finishColumn; c++ {
			la := Location{r, c}
			cell, ok := b.Map[la]
			if !ok {
				cell = &Cell{
					amountAdjacentBlackHole: 1,
				}
				b.Map[la] = cell
				continue
			}
			cell.amountAdjacentBlackHole++
		}
	}
}
func (b Board) showMap() {
	for r := 0; r < b.Rows; r++ {
		row := ""
		for c := 0; c < b.Columns; c++ {
			cell, ok := b.Map[Location{r, c}]
			if !ok {
				row += " 0 "
				continue
			}
			if cell.blackHole {
				row += " * "
				continue
			}
			row += fmt.Sprintf(" %d ", cell.amountAdjacentBlackHole)
		}
		fmt.Println(row)
	}
}