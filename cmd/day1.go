package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day1Cmd)
}

var day1Cmd = &cobra.Command{
	Use: "day1",
	Run: func(cmd *cobra.Command, args []string) {
		day1()
	},
}

func day1() {
	input, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer input.Close()

	var a, b []int
	regex := regexp.MustCompile(`(\d+)\s+(\d+)`)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			r1, err := strconv.Atoi(match[1])
			if err != nil {
				panic(fmt.Errorf("invalid number in input: %w", err))
			}
			r2, err := strconv.Atoi(match[2])
			if err != nil {
				panic(fmt.Errorf("invalid number in input: %w", err))
			}
			a = append(a, r1)
			b = append(b, r2)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading input file: %w", err))
	}

	sort.Ints(a)
	sort.Ints(b)

	dif := 0
	for i := 0; i < len(a); i++ {
		dif += int(math.Abs(float64(a[i] - b[i])))
	}

	// part 1:  2375403
	fmt.Println("part 1: ", dif)

	dif = 0
	for i := 0; i < len(a); i++ {
		count := 0
		for i2 := 0; i2 < len(a); i2++ {
			if a[i] == b[i2] {
				count++
			}
		}
		if count > 0 {
			dif += a[i] * count
		}
	}

	// part 2:  23082277
	fmt.Println("part 2: ", dif)
}
