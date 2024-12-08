package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	OPERATORS = []string{"+", "*", "||"}
)

func main() {
	part := os.Args[1]
	if part == "" || part == "1" {
		part1()
		return
	}

	part2()
}

func part1() {
	equations := getInput()

	sum := 0
	for k, v := range equations {
		// fmt.Printf("%d: %+v\n", k, v)
		if TestOperators(k, v) {
			sum += k
		}
	}

	fmt.Println(sum)

}

func part2() {
}

func getInput() map[int][]int {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	contentstr := string(content)
	lines := strings.Split(contentstr, "\n")
	input := map[int][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		eq := strings.Split(line, ":")
		ans, err := strconv.Atoi(eq[0])
		if err != nil {
			panic(err)
		}

		eq[1] = strings.TrimSpace(eq[1])
		vals := strings.Split(eq[1], " ")

		intVals := []int{}
		for _, v := range vals {
			if v == "" {
				continue
			}
			intVal, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			intVals = append(intVals, intVal)
		}

		input[ans] = intVals

	}

	return input
}

func TestOperators(ans int, vals []int) bool {
	combinations := [][]string{}
	combinations = addCombinations(combinations, 0, len(vals)-1)

	for _, combo := range combinations {
		val := vals[0]
		for i, op := range combo {
			switch op {
			case "+":
				val += vals[i+1]
			case "*":
				val *= vals[i+1]
			case "||":
				cVal, err := combineVals(val, vals[i+1])
				if err != nil {
					panic(err)
				}
				val = cVal
			}
		}
		if val == ans {
			return true
		}
	}
	return false
}

func addCombinations(combinations [][]string, i, n int) [][]string {
	if i == n {
		return combinations
	}

	if len(combinations) == 0 {
		for _, op := range OPERATORS {
			combinations = append(combinations, []string{op})
		}
		return addCombinations(combinations, i+1, n)
	}
	duplicateCombos := [][]string{}
	for j := range combinations {
		tmpCombo := []string{}
		tmpCombo = append(tmpCombo, combinations[j]...)
		duplicateCombos = append(duplicateCombos, tmpCombo)
	}

	for i, op := range OPERATORS {
		duplicate := [][]string{}
		for j := range duplicateCombos {
			tmpCombo := []string{}
			tmpCombo = append(tmpCombo, duplicateCombos[j]...)
			duplicate = append(duplicate, tmpCombo)
		}
		// fmt.Printf("%d- %+v\n", i, combinations)
		// fmt.Printf("%d- %+v\n", i, duplicate)

		for j := range duplicateCombos {
			duplicate[j] = append(duplicate[j], op)
		}

		if i == 0 {
			combinations = duplicate
		} else {
			combinations = append(combinations, duplicate...)
		}
	}

	return addCombinations(combinations, i+1, n)
}

func combineVals(a, b int) (int, error) {
	aStr := strconv.Itoa(a)
	aStr += strconv.Itoa(b)

	return strconv.Atoi(aStr)
}
