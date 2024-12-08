package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "day6",
			Run: func(cmd *cobra.Command, args []string) {
				day6()
			},
		},
	)
}

func day6() {
	input, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()

		var c []rune
		for _, r := range line {
			c = append(c, r)
		}

		grid = append(grid, c)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	startX, startY := 0, 0
	GUARD := int32('^')
	OBSTRUCTION := int32('#')
	VISITED := int32('V')
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == GUARD {
				startY = y
				startX = x
				break
			}
		}
	}

	part1 := 0
	rotation := 0
	dir := [][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	posX, posY := startX, startY
	for {
		nextX := posX + dir[rotation][1]
		nextY := posY + dir[rotation][0]
		if !withinBounds(grid, nextY, nextX) {
			part1++
			break
		}
		if grid[nextY][nextX] == OBSTRUCTION {
			rotation = (rotation + 1) % 4
			continue
		}
		if grid[posY][posX] != VISITED {
			grid[posY][posX] = VISITED
			part1++
		}
		posX, posY = nextX, nextY
	}
	fmt.Println("part1: ", part1) // 5564

	// todo revisit this solution, must be a better way to see if we're stuck in a loop
	part2 := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {

			if grid[y][x] == OBSTRUCTION {
				continue
			}

			copy := grid[y][x]
			grid[y][x] = OBSTRUCTION

			rotation = 0
			posX, posY = startX, startY

			steps := 0

			for steps < 10_000 {
				nextX := posX + dir[rotation][1]
				nextY := posY + dir[rotation][0]
				if !withinBounds(grid, nextY, nextX) {
					break
				}
				if grid[nextY][nextX] == OBSTRUCTION {
					rotation = (rotation + 1) % 4
					continue
				}
				if grid[posY][posX] != VISITED {
					grid[posY][posX] = VISITED
				}
				posX, posY = nextX, nextY
				steps++
			}

			if steps == 10_000 {
				part2++
			}

			grid[y][x] = copy
		}
	}

	fmt.Println("part2: ", part2) // 1976
}
