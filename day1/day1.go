package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	leftColumn, rightColumn := []int{}, []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		leftNumber, _ := strconv.Atoi(parts[0])
		rightNumber, _ := strconv.Atoi(parts[1])
		leftColumn = append(leftColumn, leftNumber)
		rightColumn = append(rightColumn, rightNumber)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	answer := 0

	for index, _ := range leftColumn {
		temp := rightColumn[index] - leftColumn[index]
		if temp < 0 {
			temp = temp * -1
		}

		answer += temp
	}
	return answer

}

func Part2() int {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	leftColumn, rightColumn := []int{}, []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		leftNumber, _ := strconv.Atoi(parts[0])
		rightNumber, _ := strconv.Atoi(parts[1])
		leftColumn = append(leftColumn, leftNumber)
		rightColumn = append(rightColumn, rightNumber)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	intMap := make(map[int]int)
	answer := 0
	count := 0

	for _, v := range leftColumn {
		if slices.Contains(rightColumn, v) {

			for _, innerV := range rightColumn {
				if innerV == v {
					count++
				}
			}
			intMap[v] = count
			count = 0

		}
	}

	for key, value := range intMap {
		answer += key * value
	}
	return answer

}
