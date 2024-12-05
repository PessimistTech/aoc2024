package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	rules, updates := getInput()

	// for _, rule := range rules {
	// 	fmt.Printf("%d | %d\n", rule[0], rule[1])
	// }
	// for _, update := range updates {
	// 	fmt.Printf("%+v\n", update)
	// }

	correctUpdates, _ := determineCorrectUpdates(rules, updates)
	// for _, update := range correctUpdates {
	// 	fmt.Printf("%+v\n", update)
	// }

	sum := sumCenters(correctUpdates)
	fmt.Println(sum)

}

func part2() {
	rules, updates := getInput()
	_, incorrectUpdates := determineCorrectUpdates(rules, updates)
	// for _, update := range incorrectUpdates {
	// 	fmt.Printf("%+v\n", update)
	// }

	for i, update := range incorrectUpdates {
		incorrectUpdates[i] = fixRules(rules, update)
	}
	// for _, update := range incorrectUpdates {
	// 	fmt.Printf("%+v\n", update)
	// }

	sum := sumCenters(incorrectUpdates)
	fmt.Println(sum)
}

func getInput() ([][2]int, [][]int) {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	rules := [][2]int{}
	updates := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			page1, err := strconv.Atoi(pages[0])
			if err != nil {
				panic(err)
			}
			page2, err := strconv.Atoi(pages[1])
			if err != nil {
				panic(err)
			}

			rule := [2]int{page1, page2}
			rules = append(rules, rule)
			continue
		}

		updateStrPages := strings.Split(line, ",")
		updatePages := []int{}
		for _, page := range updateStrPages {
			pageInt, err := strconv.Atoi(page)
			if err != nil {
				panic(err)
			}

			updatePages = append(updatePages, pageInt)
		}
		updates = append(updates, updatePages)
	}

	return rules, updates
}

func determineCorrectUpdates(rules [][2]int, updates [][]int) ([][]int, [][]int) {
	correctUpdates := [][]int{}
	incorrectUpdates := [][]int{}
	for _, update := range updates {
		valid := true
		for _, rule := range rules {
			if !checkRule(rule, update) {
				fmt.Printf("rule failed. Rule: %d|%d Update: %+v\n", rule[0], rule[1], update)
				valid = false
				break
			}
		}

		if valid {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	return correctUpdates, incorrectUpdates
}

func checkRule(rule [2]int, update []int) bool {
	idx1 := -1
	idx2 := -1
	for i, page := range update {
		if page == rule[0] {
			idx1 = i
			if idx2 == -1 {
				return true
			}
		}

		if page == rule[1] {
			idx2 = i
		}
	}
	if idx2 < idx1 {
		return false
	}

	return true
}

func sumCenters(updates [][]int) int {
	sum := 0
	for _, update := range updates {
		center := math.Floor(float64(len(update)) / 2)
		fmt.Printf("Adding %d (idx: %d)\n", update[int(center)], int(center))
		sum += update[int(center)]
	}

	return sum
}

func fixRules(rules [][2]int, update []int) []int {
	for _, rule := range rules {
		update = fixRule(rule, update)
	}

	if !checkUpdate(rules, update) {
		return fixRules(rules, update)
	}

	return update
}

func checkUpdate(rules [][2]int, update []int) bool {
	for _, rule := range rules {
		if !checkRule(rule, update) {
			return false
		}
	}

	return true
}

func fixRule(rule [2]int, update []int) []int {
	idx1 := -1
	idx2 := -1
	for i, page := range update {
		if page == rule[0] {
			idx1 = i
		}

		if page == rule[1] {
			idx2 = i
		}
	}
	if idx2 < idx1 && idx2 != -1 && idx1 != -1 {
		fmt.Printf("Starting: %+v\n", update)
		tmp := update[idx1]
		update[idx1] = update[idx2]
		update[idx2] = tmp
		fmt.Printf("After: %+v\n", update)
	}

	return update
}
