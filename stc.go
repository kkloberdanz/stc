// Copyright (c) 2021 Kyle Kloberdanz

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	firstError := true
	firstConversion := true

	reader := bufio.NewReader(os.Stdin)
	lines := 0
	sum := 0.0
	max := 0.0
	min := 0.0

	var allNumbers []float64

	aPtr := flag.Bool(
		"a",
		false,
		"enables statistics that require memory allocation",
	)
	flag.Parse()

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		text = strings.ReplaceAll(text, "\n", "")
		asF64, err := strconv.ParseFloat(text, 64)
		if err != nil {
			msg := "failed to convert to a number, ignoring future warnings.\n"
			if firstError {
				fmt.Fprintf(
					os.Stderr,
					msg,
				)
			}
			firstError = false
			continue
		}

		if firstConversion {
			max = asF64
			min = asF64
		}
		firstConversion = false

		lines++
		sum += asF64
		max = math.Max(max, asF64)
		min = math.Min(min, asF64)

		if *aPtr {
			allNumbers = append(allNumbers, asF64)
		}
	}

	if lines == 0 {
		fmt.Fprintf(os.Stderr, "no lines in input\n")
		os.Exit(1)
	}

	mean := sum / float64(lines)

	var median float64 = 0.0
	var mode float64 = 0.0
	var pct1 float64 = 0.0
	var pct5 float64 = 0.0
	var pct10 float64 = 0.0
	var pct25 float64 = 0.0
	var pct75 float64 = 0.0
	var pct95 float64 = 0.0
	var pct99 float64 = 0.0
	var pct99_9 float64 = 0.0
	var variance float64 = 0.0
	modeCount := 0

	if *aPtr {
		sort.Float64s(allNumbers)
		allNumbersLen := len(allNumbers)
		if allNumbersLen%2 == 0 {
			a := allNumbers[(allNumbersLen/2)-1]
			b := allNumbers[allNumbersLen/2]
			median = (a + b) / 2
		} else {
			median = allNumbers[allNumbersLen/2]
		}

		floatAllNumbersLen := float64(allNumbersLen)
		pct1 = allNumbers[int64(floatAllNumbersLen*0.01)]
		pct5 = allNumbers[int64(floatAllNumbersLen*0.05)]
		pct10 = allNumbers[int64(floatAllNumbersLen*0.1)]
		pct25 = allNumbers[int64(floatAllNumbersLen*0.25)]
		pct75 = allNumbers[int64(floatAllNumbersLen*0.75)]
		pct95 = allNumbers[int64(floatAllNumbersLen*0.95)]
		pct99 = allNumbers[int64(floatAllNumbersLen*0.99)]
		pct99_9 = allNumbers[int64(floatAllNumbersLen*0.999)]

		currentCount := 0
		prevNum := 0.0
		candidateMode := 0.0
		for i, x := range allNumbers {
			variance += (1.0 / float64(lines)) * (x - mean) * (x - mean)

			if i == 0 {
				candidateMode = x
				prevNum = x
				currentCount++
			} else if x == prevNum {
				candidateMode = x
				currentCount++
			} else {
				if currentCount > modeCount {
					mode = candidateMode
					modeCount = currentCount
				}
				prevNum = x
				currentCount = 1
			}
		}
	}

	if *aPtr {
		fmt.Printf(
			"lines: %d sum: %g mean: %g max: %g min: %g median: %g "+
				"mode: (%g %dx) variance: %g stddev: %g "+
				"pct1: %g pct5: %g pct10: %g pct25: %g pct75: %g pct95: %g "+
				"pct99: %g pct99.9: %g\n",
			lines,
			sum,
			mean,
			max,
			min,
			median,
			mode,
			modeCount,
			variance,
			math.Sqrt(variance),
			pct1,
			pct5,
			pct10,
			pct25,
			pct75,
			pct95,
			pct99,
			pct99_9,
		)
	} else {
		fmt.Printf(
			"lines: %d sum: %g mean: %g max: %g min: %g\n",
			lines,
			sum,
			mean,
			max,
			min,
		)
	}
}
