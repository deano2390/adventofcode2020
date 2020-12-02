package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("expensereport.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int64

	for scanner.Scan() {

		if i, err := strconv.ParseInt(scanner.Text(), 10, 64); err == nil {
			numbers = append(numbers, i)
		}

	}

	file.Close()

	var part1 int64 = part1(numbers)
	fmt.Printf("part1 answer: %d \n", part1)

	var part2 int64 = part2(numbers)
	fmt.Printf("part2 answer: %d \n", part2)

	part1 = findSolutionUsingRecursion(numbers, 2)
	fmt.Printf("part1 answer using recursion: %d \n", part1)

	part1 = findSolutionUsingRecursion(numbers, 3)
	fmt.Printf("part2 answer using recursion: %d \n", part2)

}

func part1(numbers []int64) int64 {
	for _, outterLoop := range numbers {
		for _, innerLoop := range numbers {

			if outterLoop+innerLoop == 2020 {
				fmt.Printf("solution found: %d & %d \n", outterLoop, innerLoop)
				return outterLoop * innerLoop
			}
		}
	}

	return 0
}

func part2(numbers []int64) int64 {
	for _, loop1 := range numbers {
		for _, loop2 := range numbers {
			for _, loop3 := range numbers {
				if loop1+loop2+loop3 == 2020 {
					fmt.Printf("solution found: %d & %d & %d\n", loop1, loop2, loop3)
					return loop1 * loop2 * loop3
				}
			}
		}
	}

	return 0
}

func findSolutionUsingRecursion(numbers []int64, requiredDepth int64) int64 {
	return recurse(2020, 0, numbers, 2, 0)
}

func recurse(targetSum int64, currentSum int64, numbers []int64, requiredDepth int64, currentDepth int64) int64 {

	currentDepth++

	for _, number := range numbers {
		var newSum int64 = currentSum + number
		if currentDepth == requiredDepth {
			if newSum == targetSum {
				fmt.Printf("number found: %d\n", number)
				return number
			}
		} else {
			var result int64 = recurse(targetSum, newSum, numbers, requiredDepth, currentDepth)
			if result > -1 {
				fmt.Printf("number found: %d\n", number)
				return result * number
			}
		}
	}

	return -1
}
