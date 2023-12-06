package aoc23

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const prodInput = `62     73     75     65
644   1023   1240   1023`

const testInput1 = `7  15   30
9  40  200`

var inputs = [][2]string{
	{"testInput1", testInput1},
	{"prodInput", prodInput},
}

func part1(t *testing.T, input string) {
	reader := bufio.NewScanner(strings.NewReader(input))
	r := must(t, regexp.Compile, `\d+`)
	var lineNums [][]int
	for reader.Scan() {
		line := reader.Text()
		nums := mapSlice(func(in string) int {
			return must(t, strconv.Atoi, in)
		}, r.FindAllString(line, -1))
		lineNums = append(lineNums, nums)
	}
	require.NoError(t, reader.Err())
	fmt.Println(lineNums)

	answer := 1
	for raceIdx := range lineNums[0] {
		winningMethods := 0
		raceDuration := lineNums[0][raceIdx]
		//recordDistance := lineNums[1][i]
		for chargeDuration := 0; chargeDuration < raceDuration; chargeDuration++ {
			speed := chargeDuration
			distanceCovered := speed * (raceDuration - chargeDuration)
			if distanceCovered > lineNums[1][raceIdx] {
				winningMethods++
			}
		}
		fmt.Println(winningMethods)
		answer *= winningMethods
	}
	fmt.Println(answer)
}

func mapSlice[in any, out any](mapIt func(in) out, inSlice []in) []out {
	outSlice := make([]out, len(inSlice))
	for i, inV := range inSlice {
		outSlice[i] = mapIt(inV)
	}
	return outSlice
}

func must[I any, O any](t *testing.T, fn func(I) (O, error), in I) O {
	o, err := fn(in)
	require.NoError(t, err)
	return o
}

func TestPart1(t *testing.T) {
	for _, input := range inputs {
		fmt.Println(input[0])
		part1(t, input[1])
		fmt.Println()
	}
}
