package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func day2() {
	input, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	regex := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(input)

	var lines [][]int

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllString(line, -1)

		var n []int
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			n = append(n, num)
		}

		lines = append(lines, n)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	safe := 0
	for i := 0; i < len(lines); i++ {
		if isSafe(lines[i]) {
			safe++
		}
	}

	println("part1: ", safe) // 631

	safe = 0
	for i := 0; i < len(lines); i++ {
		if isSafe(lines[i]) {
			safe++
			continue
		}

		for i2 := 0; i2 < len(lines[i]); i2++ {
			if isSafe(remove(lines[i], i2)) {
				safe++
				break
			}
		}
	}

	println("part2: ", safe) // 665
}

func remove(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)

	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])

	return newSlice
}

func isSafe(line []int) bool {
	last := line[0]
	increase := line[1] > line[0]
	for i := 1; i < len(line); i++ {
		n := line[i]
		if increase && n < last || !increase && n > last {
			return false
		}
		d := int(math.Abs(float64(last - n)))
		if d < 1 || d > 3 {
			return false
		}
		last = n
	}
	return true
}
