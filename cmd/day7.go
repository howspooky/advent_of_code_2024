package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "day7",
			Run: func(cmd *cobra.Command, args []string) {
				day7()
			},
		},
	)
}

type equation struct {
	sum     int
	numbers []int
}

func day7() {
	input, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	regex := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(input)

	var equations []equation

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllString(line, -1)

		sum, err := strconv.Atoi(matches[0])
		if err != nil {
			panic(err)
		}

		var n []int
		for i := 1; i < len(matches); i++ {
			number, err := strconv.Atoi(matches[i])
			if err != nil {
				panic(err)
			}
			n = append(n, number)
		}

		e := equation{
			sum:     sum,
			numbers: n,
		}

		equations = append(equations, e)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	part1 := 0
	for _, e := range equations {
		current := []int{e.numbers[0]}

		for _, num := range e.numbers[1:] {
			var next []int
			for _, cur := range current {
				next = append(next, cur*num)
				next = append(next, cur+num)
			}
			current = next
		}

		if slices.Contains(current, e.sum) {
			part1 += e.sum
		}
	}
	fmt.Println("part1: ", part1) // 4555081946288

	part2 := 0
	for _, e := range equations {
		current := []int{e.numbers[0]}

		for _, num := range e.numbers[1:] {
			var next []int
			for _, cur := range current {
				concatenated := cur
				temp := num
				for temp > 0 {
					concatenated *= 10
					temp /= 10
				}
				concatenated += num

				next = append(next, cur+num)
				next = append(next, cur*num)
				next = append(next, concatenated)
			}
			current = next
		}

		if slices.Contains(current, e.sum) {
			part2 += e.sum
		}
	}

	fmt.Println("part2: ", part2) // 227921760109726
}
