package aoc23

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"regexp"
	"sort"
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
	var lineNums [][]string
	for reader.Scan() {
		line := reader.Text()
		lineNums = append(lineNums, r.FindAllString(line, -1))
	}
	raceDuration := must(t, func(i string) (int64, error) {
		return strconv.ParseInt(i, 10, 64)
	}, strings.Join(lineNums[0], ""))
	raceRecord := must(t, func(i string) (int64, error) {
		return strconv.ParseInt(i, 10, 64)
	}, strings.Join(lineNums[1], ""))
	require.NoError(t, reader.Err())
	fmt.Println(raceDuration, raceRecord)

	// speed = durationCharging
	// dist = speed * (duration - durationCharging)
	// dist = durationCharging * (duration - durationCharging)
	// dist = durationCharging * duration - durationCharging^2
	// durationCharging^2 - duration * durationCharging + dist = 0

	// to get durationCharging for 0
	// ax^2 + bx + c = 0
	// a = 1
	// b = -duration
	// c = dist
	// quadratic
	// discriminant := b*b - 4.0*c
	b := -raceDuration
	c := raceRecord
	discriminant := b*b - 4*c
	d := math.Sqrt(float64(discriminant))
	answer1 := (float64(-b) + d) / 2.0
	fmt.Printf("%0.6f\n", answer1)
	answer2 := (float64(-b) - d) / 2.0
	fmt.Printf("%0.6f\n", answer2)

	answers := []float64{answer1, answer2}
	sort.Float64s(answers)
	fmt.Println(answers)
	fmt.Println("lower")
	fmt.Println(math.Floor(answers[0] + 1))
	fmt.Println(fmt.Sprintf("%0.20f", math.Floor(answers[0]+1)))
	fmt.Println(math.Ceil(answers[0]))
	fmt.Println(fmt.Sprintf("%0.20f", math.Ceil(answers[0])))

	fmt.Println("upper")
	fmt.Println(math.Ceil(answers[1] - 1))
	fmt.Println(fmt.Sprintf("%0.20f", math.Ceil(answers[1]-1)))
	fmt.Println(math.Floor(answers[1]))
	fmt.Println(fmt.Sprintf("%0.20f", math.Floor(answers[1])))

	fmt.Println("answer?")
	fmt.Println((math.Ceil(answers[1] - 1)) - math.Floor(answers[0]+1) + 1)

	// tried
	// 60649495
	// 36872656

	// distanceCovered = (raceDuration - timeCharged) * timeCharged
	// 0 = (raceDuration - timeCharged) * timeCharged
	// 0 = raceDuration*timeCharged - timeCharged^2
	// timeCharged^2 = raceDuration*timeCharged
	//

	//answer := 1
	//for raceIdx := range lineNums[0] {
	//	winningMethods := 0
	//	raceDuration := lineNums[0][raceIdx]
	//	//recordDistance := lineNums[1][i]
	//	for chargeDuration := 0; chargeDuration < raceDuration; chargeDuration++ {
	//		speed := chargeDuration
	//		distanceCovered := speed * (raceDuration - chargeDuration)
	//		if distanceCovered > lineNums[1][raceIdx] {
	//			winningMethods++
	//		}
	//	}
	//	fmt.Println(winningMethods)
	//	answer *= winningMethods
	//}
	//fmt.Println(answer)
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
