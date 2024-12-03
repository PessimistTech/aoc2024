package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("\nPart 2:")
	part2()
}

func part1() {
	list1, list2, size := getInput()
	slices.Sort(list1)
	slices.Sort(list2)

	result := []int{}
	for i := 0; i < size; i++ {
		res := 0
		if list1[i] > list2[i] {
			res = list1[i] - list2[i]
		} else {
			res = list2[i] - list1[i]
		}
		result = append(result, res)
	}

	sum := 0
	for _, val := range result {
		sum += val
	}

	fmt.Println(sum)

}

func part2() {
	list1, list2, _ := getInput()

	similarityScores := []int{}
	for _, r := range list1 {
		count := 0
		for _, c := range list2 {
			if r == c {
				count++
			}
		}
		similarityScores = append(similarityScores, r*count)
	}

	sum := 0
	for _, val := range similarityScores {
		sum += val
	}

	fmt.Println(sum)

}

func getInput() ([]int, []int, int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	list1 := []int{}
	list2 := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		elements := strings.Split(line, "   ")
		if len(elements) != 2 {
			panic(fmt.Sprintf("not the right length of elements in line. Got %d", len(elements)))
		}
		e1, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(err)
		}
		e2, err := strconv.Atoi(elements[1])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, e1)
		list2 = append(list2, e2)
	}

	return list1, list2, len(lines) - 1
}
