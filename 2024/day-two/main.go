package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger := log.Default()

	f, err := os.Open("./input.txt")
	if err != nil {
		logger.Fatalf("An error occurred opening the file %v", err)
	}

	r := bufio.NewReader(f)

	/**
	read the lines one by one
	for each line check if the numbers are all increasing or all decreasing
	and check if the difference between adjacent levels are at least 1 or at most 3
	increment he count of safe reports
	*/

	var safeReports int
	for {
		report, err := r.ReadString('\n')
		if err != nil {
			break
		}

		report = strings.TrimRight(report, "\n")

		levels := strings.Split(report, " ")

		meetsAdjCondition := false
		strictlyInc := false
		strictlyDec := false
		for idx := 0; idx < len(levels)-1; idx++ {
			fl, _ := strconv.Atoi(levels[idx])
			sl, _ := strconv.Atoi(levels[idx+1])

			// check if the numbers are strictly increasing
			// or if they are strictly decreasing
			// if either one, don't do the other,
			// if one fails, check the other

			if fl > sl {
				strictlyInc = true
			} else if fl < sl {
				strictlyDec = true
			}

			// check the difference between adjacent neighbors
			nd := math.Abs(float64(fl - sl))
			if nd >= 1 && nd <= 3 {
				meetsAdjCondition = true
			} else {
				meetsAdjCondition = false
				break
			}

		}

		if meetsAdjCondition && (!(strictlyInc && strictlyDec)) {
			safeReports += 1
		}

		logger.Println(levels, safeReports)
	}
}
