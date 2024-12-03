package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("An error occurred opening the file %v", err)
	}

	r := bufio.NewReader(f)

	var listA []int
	var listB []int
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			break
		}

		nums := strings.Split(line, ",")

		num1, _ := strconv.Atoi(nums[0])
		listA = append(listA, num1)

		num2, err := strconv.Atoi(strings.TrimRight(nums[1], "\n"))
		listB = append(listB, num2)

	}

	slices.SortFunc(listA, func(a, b int) int {
		return a - b
	})

	slices.SortFunc(listB, func(a, b int) int {
		return a - b
	})

	var totalDiff int
	for idx, num := range listA {
		totalDiff = totalDiff + (int(math.Abs(float64(num - listB[idx]))))
	}

	fmt.Println("Total Diff ", totalDiff)

	var similarityScore int
	for _, num := range listA {
		var numOfAppearance int
		for _, snum := range listB {
			if num == snum {
				numOfAppearance += 1
			}
		}

		similarityScore = similarityScore + (num * numOfAppearance)
	}

	fmt.Println("Similarity Score ", similarityScore)

}
