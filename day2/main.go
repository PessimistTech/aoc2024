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
	reports, err := getInput()
	if err != nil {
		panic(err)
	}

	safeCount := 0
	for _, report := range reports {
		if checkReportSafety(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part2() {
	reports, err := getInput()
	if err != nil {
		panic(err)
	}

	safeCount := 0
	for _, report := range reports {
		if checkReportSafety(report) {
			safeCount++
			continue
		}

		if applyDampener(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func getInput() ([][]int, error) {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		return [][]int{}, err
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	reports := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		strValues := strings.Split(line, " ")

		values, err := convertToInt(strValues)
		if err != nil {
			return reports, err
		}

		reports = append(reports, values)
	}

	return reports, nil
}

func convertToInt(strIn []string) ([]int, error) {
	result := []int{}

	for _, str := range strIn {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			return result, err
		}

		result = append(result, intVal)
	}

	return result, nil
}

func checkReportSafety(report []int) bool {
	inDec, startDiff := compare(report[0], report[1])
	if startDiff > 3 || startDiff < 1 {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		change, diff := compare(report[i], report[i+1])

		if change != inDec {
			return false
		}

		if diff > 3 || diff < 1 {
			return false
		}

	}

	return true
}

func compare(a, b int) (string, float64) {
	diff := a - b
	dist := math.Abs(float64(diff))
	if dist == 0 {
		return "equal", dist
	}
	if diff < 0 {
		return "incrase", dist
	}

	if diff > 0 {
		return "decrease", dist
	}

	return "", dist
}

func applyDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		tmpReport := removeLevel(i, report)
		if checkReportSafety(tmpReport) {
			return true
		}
	}

	return false
}

func removeLevel(id int, report []int) []int {
	newReport := []int{}
	for i, level := range report {
		if i == id {
			continue
		}

		newReport = append(newReport, level)
	}

	return newReport
}
