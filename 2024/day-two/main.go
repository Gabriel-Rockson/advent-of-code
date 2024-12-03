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
	defer f.Close()

	r := bufio.NewReader(f)

	var safeReports int
	for {
		report, err := r.ReadString('\n')
		if err != nil {
			break
		}

		report = strings.TrimRight(report, "\n")
		levels := strings.Split(report, " ")

		valid := checkSequence(levels)

		if !valid {
			for i := 0; i < len(levels); i++ {
				newLevels := make([]string, 0)
				newLevels = append(newLevels, levels[:i]...)
				newLevels = append(newLevels, levels[i+1:]...)

				if checkSequence(newLevels) {
					valid = true
					break
				}
			}
		}

		if valid {
			safeReports++
		}
	}

	logger.Printf("Total Safe Reports: %d\n", safeReports)
}

func checkSequence(levels []string) bool {
	if len(levels) < 2 {
		return false
	}

	strictlyInc := false
	strictlyDec := false

	for idx := 0; idx < len(levels)-1; idx++ {
		fl, err1 := strconv.Atoi(levels[idx])
		sl, err2 := strconv.Atoi(levels[idx+1])

		if err1 != nil || err2 != nil {
			return false
		}

		if fl > sl {
			if strictlyDec {
				return false
			}
			strictlyInc = true
		} else if fl < sl {
			if strictlyInc {
				return false
			}
			strictlyDec = true
		}

		nd := math.Abs(float64(fl - sl))
		if nd < 1 || nd > 3 {
			return false
		}
	}

	return !(strictlyInc && strictlyDec)
}
