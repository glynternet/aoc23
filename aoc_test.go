package aoc23

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const prodInput = ``

const testInput = ``

var inputs = [][2]string{
	{"testInput", testInput},
	{"prodInput", prodInput},
}

func part(t *testing.T, input string) {
	var directions []bool
	nodeDirs := make(map[string][2]string)
	var currents []string

	scanLines(t, input, skipBlankLines, func(line string) bool {
		if len(directions) == 0 {
			for _, char := range line {
				directions = append(directions, char == 'R')
			}
			return true
		}
		split := strings.Split(line, " ")
		node := split[0]
		if _, ok := nodeDirs[node]; ok {
			panic("here")
		}
		nodeDirs[node] = [2]string{split[2][1 : len(split[2])-1], split[3][:len(split[3])-1]}
		if node[len(node)-1] == 'A' {
			currents = append(currents, node)
		}
		return true
	})

	dirIdx := 0
	moves := 0
	firstZs := make([]int, len(currents))
	for i := range firstZs {
		firstZs[i] = -1
	}
	for ; !all(currents, func(s string) bool { return s[len(s)-1] == 'Z' }) && !all(firstZs, func(v int) bool { return v != -1 }); moves++ {
		right := directions[dirIdx%len(directions)]
		dirIdx++
		if right {
			for i := range currents {
				currents[i] = nodeDirs[currents[i]][1]
			}
		} else {
			for i := range currents {
				currents[i] = nodeDirs[currents[i]][0]
			}
		}
		for i, current := range currents {
			if firstZs[i] != -1 {
				continue
			}
			if current[len(current)-1] == 'Z' {
				firstZs[i] = moves + 1
			}
		}
	}
	fmt.Println(moves)
	fmt.Println("Find the least common multiple of these values", firstZs)
}

func TestPart1(t *testing.T) {
	for _, input := range inputs {
		fmt.Println(input[0])
		part(t, input[1])
		fmt.Println()
	}
}

func scanLines(t *testing.T, input string, processors ...func(string) bool) {
	reader := bufio.NewScanner(strings.NewReader(input))
	for reader.Scan() {
		line := reader.Text()
		fmt.Println(line)
		var processed bool
		for _, p := range processors {
			if p(line) {
				processed = true
				break
			}
		}
		require.True(t, processed, "not all input was processed for line:", line)
	}
}

func skipBlankLines(line string) bool {
	return line == ""
}

func all[t any](values []t, predicate func(t) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func mapSlice[in any, out any](mapIt func(in) out, inSlice []in) []out {
	outSlice := make([]out, len(inSlice))
	for i, inV := range inSlice {
		outSlice[i] = mapIt(inV)
	}
	return outSlice
}

func must[I any, O any](t *testing.T, fn func(I) (O, error), in I) O {
	t.Helper()
	o, err := fn(in)
	require.NoError(t, err)
	return o
}
