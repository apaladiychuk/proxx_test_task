package main

import (
	"fmt"
	"math/rand"
	"time"
)

const attemptKoef = 10

// Location struct that represents cell position on board
type Location struct {
	Row    int
	Column int
}

// Cell struct that represents information about cell
type Cell struct {
	blackHole               bool
	amountAdjacentBlackHole int
	isOpened                bool
}

// Board struct that represents infoirmation about board
type Board struct {
	Rows       int
	Columns    int
	BlackHoles int
	Map        map[Location]*Cell
}

// NewBoard returns new implementation of Board by given parameters
// rows - amount of rows in board
// columns - amount columns in board
// blackHoles amount blackholes on board
// returns error if amount of blackhole greater than board cells
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
	attempts := blackHoles * attemptKoef
	for i := 0; i < blackHoles; i++ {
		if attempts == 0 {
			return nil, fmt.Errorf("can not create board with given amount of blackholes")
		}
		loc := Location{
			Row:    int(rnd.Float64() * float64(rows)),
			Column: int(rnd.Float64() * float64(rows)),
		}
		if cell, ok := b.Map[loc]; ok && cell.blackHole {
			i--
			attempts--
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

func (b *Board) openBoard(row, column int) (bool, error) {
	if row >= b.Rows || row < 0 {
		return false, fmt.Errorf("row index out of range ")
	}
	if column >= b.Columns || column < 0 {
		return false, fmt.Errorf("column index out of range ")
	}
	loc := Location{row, column}
	cell, ok := b.Map[loc]
	if !ok {
		cell = &Cell{amountAdjacentBlackHole: 0}
	}
	if cell.isOpened {
		return true, nil
	}
	if cell.blackHole {
		cell.isOpened = true
		return false, nil
	}
	if cell.amountAdjacentBlackHole > 0 {
		cell.isOpened = true
		return true, nil
	}
	b.openAdjacent(b.getClosedAdjacement(loc))
	return true, nil
}

func (b *Board) getClosedAdjacement(loc Location) []Location {
	startColumn, finishColumn := b.adjacentIndexes(loc.Column, b.Columns)
	startRow, finishRow := b.adjacentIndexes(loc.Row, b.Rows)
	var result []Location
	for r := startRow; r <= finishRow; r++ {
		for c := startColumn; c <= finishColumn; c++ {
			la := Location{r, c}
			cell, ok := b.Map[la]
			if !ok {
				cell = &Cell{}
				b.Map[la] = cell
			}
			if !cell.isOpened {
				result = append(result, la)
			}
		}
	}
	return result
}

func (b *Board) openAdjacent(locs []Location) {
	for _, loc := range locs {
		cell, ok := b.Map[loc]
		if !ok {
			cell = &Cell{
				amountAdjacentBlackHole: 0,
				isOpened:                true,
			}
			b.Map[loc] = cell
			b.openAdjacent(b.getClosedAdjacement(loc))
			continue
		}
		cell.isOpened = true
		if cell.amountAdjacentBlackHole == 0 && !cell.blackHole {
			b.openAdjacent(b.getClosedAdjacement(loc))
		}
	}
}

func (b *Board) showMap() {
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
func (b *Board) showRealMap() {
	for r := 0; r < b.Rows; r++ {
		row := ""
		for c := 0; c < b.Columns; c++ {
			cell, ok := b.Map[Location{r, c}]
			if !ok {
				row += " - "
				continue
			}
			if !cell.isOpened {
				row += " - "
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
