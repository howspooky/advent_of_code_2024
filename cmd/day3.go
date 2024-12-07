package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day3Cmd)
}

var day3Cmd = &cobra.Command{
	Use: "day3",
	Run: func(cmd *cobra.Command, args []string) {
		day3()
	},
}

var MUL = 0
var DO = 1
var DONT = 2

func day3() {
	input, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	scanner := bufio.NewScanner(input)

	var instructions [][]int

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				instructions = append(instructions, []int{DO})
			} else if match[0] == "don't()" {
				instructions = append(instructions, []int{DONT})
			} else {
				num1, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				num2, err := strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}
				instructions = append(instructions, []int{MUL, num1, num2})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	part1 := 0
	for i := 0; i < len(instructions); i++ {
		instr := instructions[i]
		if instr[0] == MUL {
			part1 += instr[1] * instr[2]
		}
	}
	fmt.Println("part1: ", part1) // 174960292

	enabled := true
	part2 := 0
	for i := 0; i < len(instructions); i++ {
		instr := instructions[i]
		if instr[0] == DO {
			enabled = true
		} else if instr[0] == DONT {
			enabled = false
		} else if instr[0] == MUL {
			if enabled {
				part2 += instr[1] * instr[2]
			}
		}
	}
	fmt.Println("part2: ", part2) // 56275602
}
