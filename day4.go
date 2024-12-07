package main

import (
	"bufio"
	"fmt"
	"os"
)

func day4() {
	input, err := os.Open("inputs/day4.txt")
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

	X := int32('X')
	M := int32('M')
	A := int32('A')
	S := int32('S')
	XMAS := []rune{X, M, A, S}

	part1 := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != XMAS[0] {
				continue
			}
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					posY := y + dy
					posX := x + dx
					if !withinBounds(grid, posY, posX) {
						continue
					}
					if grid[posY][posX] != XMAS[1] {
						continue
					}
					found := true
					for i := 2; i < len(XMAS); i++ {
						posY2 := y + (i * dy)
						posX2 := x + (i * dx)
						if !withinBounds(grid, posY2, posX2) {
							found = false
							break
						}
						if grid[posY2][posX2] != XMAS[i] {
							found = false
							break
						}
					}
					if found {
						part1++
					}
				}
			}
		}
	}
	println("part1: ", part1) // 2414

	corners := [][2]int{
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}
	combinations := [][]rune{
		{M, S, S, M},
		{M, M, S, S},
		{S, M, M, S},
		{S, S, M, M},
	}

	part2 := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != A {
				continue
			}

			var cornerRunes []rune
			for _, d := range corners {
				dy := y + d[0]
				dx := x + d[1]
				if !withinBounds(grid, dy, dx) {
					break
				}
				cornerRunes = append(cornerRunes, grid[dy][dx])
			}

			for _, comb := range combinations {
				if runeSlicesEqual(cornerRunes, comb) {
					for _, d := range corners {
						dy := y + d[0]
						dx := x + d[1]
						if !withinBounds(grid, dy, dx) {
							break
						}
					}
					part2++
					break
				}
			}
		}
	}
	println("part2: ", part2) // 1871
}

func withinBounds(grid [][]rune, y int, x int) bool {
	if y < 0 || y >= len(grid) {
		return false
	}
	if x < 0 || x >= len(grid[y]) {
		return false
	}
	return true
}

func runeSlicesEqual(slice1, slice2 []rune) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
