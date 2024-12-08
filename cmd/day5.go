package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "day5",
			Run: func(cmd *cobra.Command, args []string) {
				day5()
			},
		},
	)
}

func day5() {
	input, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	var rules [][]int
	var updates [][]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			rules = append(rules, splitToIntSlice(line, "|"))
		} else {
			updates = append(updates, splitToIntSlice(line, ","))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	part1 := 0
	for _, u := range updates {
		ok := true
		for i, curr := range u {
			for b := i + 1; b < len(u); b++ {
				for _, rule := range rules {
					if u[b] == rule[0] && curr == rule[1] {
						ok = false
						break
					}
				}
			}
		}
		if ok {
			part1 += u[len(u)/2]
		}
	}
	fmt.Println("part1: ", part1) // 4135

	part2 := 0
	for _, u := range updates {
		ok := true
		for i, curr := range u {
			for b := i + 1; b < len(u); b++ {
				for _, rule := range rules {
					if u[b] == rule[0] && curr == rule[1] {
						ok = false
						break
					}
				}
			}
		}
		if !ok {
			sort.Slice(u, func(i, j int) bool {
				ok := true
				for _, rule := range rules {
					if u[i] == rule[0] && u[j] == rule[1] {
						ok = false
						break
					}
				}
				return ok
			})
			part2 += u[len(u)/2]
		}
	}

	fmt.Println("part2: ", part2) // 5285
}

func splitToIntSlice(line string, delimiter string) []int {
	split := strings.Split(line, delimiter)
	var slice []int
	for _, n := range split {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		slice = append(slice, num)
	}
	return slice
}
