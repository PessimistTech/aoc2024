package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	MULCOMMAND         = regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	IS_MULCOMMAND      = regexp.MustCompile("^mul\\([0-9]{1,3},[0-9]{1,3}\\)$")
	ENABLE             = regexp.MustCompile("^do\\(\\)$")
	DISABLE            = regexp.MustCompile("^don't\\(\\)$")
	MUL_ENABLE_COMMAND = regexp.MustCompile("(mul\\([0-9]{1,3},[0-9]{1,3}\\))|(do\\(\\))|(don't\\(\\))")
	VALUES             = regexp.MustCompile("[0-9]{1,3},[0-9]{1,3}")
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
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	commands := MULCOMMAND.FindAllString(input, -1)

	sum := 0
	for _, command := range commands {
		val, err := processCommand(command)
		if err != nil {
			panic(err)
		}

		sum += val
	}

	fmt.Println(sum)
}

func part2() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	commands := MUL_ENABLE_COMMAND.FindAllString(input, -1)

	enabled := true
	sum := 0
	for _, command := range commands {
		if MULCOMMAND.MatchString(command) && enabled {
			val, err := processCommand(command)
			if err != nil {
				panic(err)
			}

			sum += val
		}

		if ENABLE.MatchString(command) {
			enabled = true
		}

		if DISABLE.MatchString(command) {
			enabled = false
		}
	}

	fmt.Println(sum)

}

func getInput() (string, error) {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func processCommand(command string) (int, error) {
	values := VALUES.FindString(command)
	if values == "" {
		return 0, fmt.Errorf("values not found in %s", command)
	}

	args := strings.Split(values, ",")
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, err
	}

	b, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, err
	}

	return a * b, nil
}
