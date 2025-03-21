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

type NumberAndIndex struct {
	num   float64
	index int64
}

type int64Set map[int64]struct{}

var exists = struct{}{}

var graphString = "\033[91m*\033[0m"

func drawGraph(
	allNumbers []NumberAndIndex,
	lines int64,
	min float64,
	max float64,
	xMax int,
	yMax int,
) {
	// TODO: don't waste time on reversing the array, just use indexing
	// from the back later
	for i, j := 0, len(allNumbers)-1; i < j; i, j = i+1, j-1 {
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	}

	xFactor := float64(lines) / float64(xMax)
	for i := 0; i < len(allNumbers); i++ {
		allNumbers[i].index = int64(float64(allNumbers[i].index) / xFactor)
	}

	yFactor := (max - min) / float64(yMax)
	yNum := max
	firstIdx := 0
	for y := 0; y < yMax+1; y++ {
		// Now that allNumbers is sorted by the y value, we will now find
		// each section of this array that is in the range of the current
		// row. Once we have found the index of the last number in this
		// row, We will add the index to the rows

		indexSet := make(int64Set)
		var spaceChar = " "
		if yNum > 0 && yNum-yFactor < 0 {
			spaceChar = "-"
			yNum = 0
		} else {
			spaceChar = " "
		}

		lowestInRow := yNum
		var lastIdx = 0
		for lastIdx = firstIdx; lastIdx < len(allNumbers); lastIdx++ {
			num := allNumbers[lastIdx].num
			if num < lowestInRow {
				for i := firstIdx; i < lastIdx; i++ {
					idx := allNumbers[i].index
					indexSet[idx] = exists
				}
				firstIdx = lastIdx
				break
			}
		}

		if y == yMax {
			lastItem := allNumbers[len(allNumbers)-1]
			indexSet[lastItem.index] = exists
		}
		if yNum >= 0 {
			fmt.Printf(" %.2e |", yNum)
		} else {
			fmt.Printf("%.2e |", yNum)
		}

		// Loop through indices set and only add a point if it is
		// in the set
		for i := 0; i < xMax+1; i++ {
			idx := int64(i)
			if _, ok := indexSet[idx]; ok {
				fmt.Print(graphString)
			} else {
				fmt.Print(spaceChar)
			}
		}
		yNum -= yFactor
		fmt.Printf("\n")
	}
}

func main() {
	firstError := true
	firstConversion := true

	reader := bufio.NewReader(os.Stdin)
	var lines int64 = 0
	sum := 0.0
	max := 0.0
	min := 0.0

	var allNumbers []NumberAndIndex

	aPtr := flag.Bool(
		"a",
		false,
		"enables statistics that require memory allocation",
	)

	gPtr := flag.Bool(
		"g",
		false,
		"graph the data in the terminal",
	)
	xDim := flag.Int("xdim", 68, "character length of x axis")
	yDim := flag.Int("ydim", 20, "character length of y axis")

	noColor := flag.Bool("nocolor", false, "don't use colors when graphing")

	flag.Parse()

	if *noColor {
		graphString = "*"
	}

	var saveData = *aPtr || *gPtr

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		text = strings.TrimSpace(text)
		asF64, err := strconv.ParseFloat(text, 64)
		if err != nil {
			msg := "failed to convert to a number, ignoring future warnings."
			if firstError {
				fmt.Fprintln(
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

		if saveData {
			num := NumberAndIndex{
				num:   asF64,
				index: lines,
			}
			allNumbers = append(allNumbers, num)
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

	if saveData {
		sort.Slice(allNumbers, func(i, j int) bool {
			return allNumbers[i].num < allNumbers[j].num
		})
		mode = allNumbers[0].num
		allNumbersLen := len(allNumbers)
		if allNumbersLen%2 == 0 {
			a := allNumbers[(allNumbersLen/2)-1].num
			b := allNumbers[allNumbersLen/2].num
			median = (a + b) / 2
		} else {
			median = allNumbers[allNumbersLen/2].num
		}

		floatAllNumbersLen := float64(allNumbersLen)
		pct1 = allNumbers[int64(floatAllNumbersLen*0.01)].num
		pct5 = allNumbers[int64(floatAllNumbersLen*0.05)].num
		pct10 = allNumbers[int64(floatAllNumbersLen*0.1)].num
		pct25 = allNumbers[int64(floatAllNumbersLen*0.25)].num
		pct75 = allNumbers[int64(floatAllNumbersLen*0.75)].num
		pct95 = allNumbers[int64(floatAllNumbersLen*0.95)].num
		pct99 = allNumbers[int64(floatAllNumbersLen*0.99)].num
		pct99_9 = allNumbers[int64(floatAllNumbersLen*0.999)].num

		currentCount := 0
		prevNum := 0.0
		candidateMode := 0.0
		nInv := 1.0 / float64(lines)
		for i, item := range allNumbers {
			x := item.num
			difference := x - mean
			variance += nInv * difference * difference

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
		if modeCount == 0 {
			modeCount = currentCount
		}
	}

	if *gPtr {
		drawGraph(allNumbers, lines, min, max, *xDim, *yDim)
	}

	if *aPtr {
		fmt.Printf(
			"count: %d sum: %g mean: %g max: %g min: %g median: %g "+
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
