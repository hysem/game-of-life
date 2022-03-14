package game

import (
	"bytes"
	"fmt"
)

type GameOfLife struct {
	universe [][]int

	generation int

	rows int
	cols int
}

func newEmpty2DArray(rows, cols int) [][]int {
	result := make([][]int, rows)
	for row := range result {
		result[row] = make([]int, cols)
	}
	return result
}

func NewGameOfLife(rows, cols int) *GameOfLife {
	return &GameOfLife{
		rows:     rows,
		cols:     cols,
		universe: newEmpty2DArray(rows, cols),
	}

}

func (g *GameOfLife) String() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Generation: %d\n", g.generation)
	for _, row := range g.universe {
		for _, v := range row {
			fmt.Fprintf(buf, "%d ", v)
		}
		fmt.Fprintln(buf)
	}
	fmt.Fprintln(buf)
	return buf.String()
}

func (g *GameOfLife) Snapshot() {
	fmt.Println(g)
}

func (g *GameOfLife) Next() {
	g.generation++

	nextUniverse := newEmpty2DArray(g.rows, g.cols)

	var (
		liveNeighbourCount int
		isLiveCell         bool
		shouldLive         int
	)
	for row := 1; row < g.rows-1; row++ {
		for col := 1; col < g.cols-1; col++ {
			liveNeighbourCount = g.universe[row-1][col-1] + g.universe[row-1][col] + g.universe[row-1][col+1] +
				g.universe[row][col-1] + g.universe[row][col+1] +
				g.universe[row+1][col-1] + g.universe[row+1][col] + g.universe[row+1][col+1]

			isLiveCell = g.universe[row][col] == 1

			shouldLive = 0
			if isLiveCell {
				if liveNeighbourCount < 2 {
					// Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
					shouldLive = 0
				} else if liveNeighbourCount == 2 || liveNeighbourCount == 3 {
					shouldLive = 1
					// Any live cell with two or three live neighbors lives on to the next generation.
				} else if liveNeighbourCount > 3 {
					// Any live cell with more than three live neighbors dies, as if by overcrowding.
					shouldLive = 0
				}
			} else if !isLiveCell && liveNeighbourCount == 3 {
				//Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				shouldLive = 1
			}
			nextUniverse[row][col] = shouldLive
		}
	}

	g.universe = nextUniverse
}

func (g *GameOfLife) Seed(row, col int) {
	g.universe[row][col] = 1
}
