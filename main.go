package main

import (
	"fmt"
	"game-of-life/game"
	"os"
)

const (
	row, col = 25, 25
)

func main() {
	g := game.NewGameOfLife(row, col)

	centerRow := row / 2
	centerCol := col / 2

	g.Seed(centerRow-1, centerCol)
	g.Seed(centerRow, centerCol+1)
	g.Seed(centerRow+1, centerCol-1)
	g.Seed(centerRow+1, centerCol)
	g.Seed(centerRow+1, centerCol+1)
	g.Snapshot()

	var choice int
	for {
		fmt.Print("1. Next\n2. End\nChoice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			g.Next()
			g.Snapshot()
		case 2:
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}
}
