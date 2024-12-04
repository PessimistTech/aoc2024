package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	SEARCHES = []string{"M", "A", "S"}
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
	search, err := getInput()
	if err != nil {
		panic(err)
	}

	// for _, line := range search {
	// 	fmt.Printf("%+v\n", line)
	// }

	count := findXmas(search)
	fmt.Println(count)
}

func part2() {
	search, err := getInput()
	if err != nil {
		panic(err)
	}

	count := findCrossmas(search)

	fmt.Println(count)

}

func getInput() ([][]string, error) {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		return [][]string{}, err
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)
	lines := strings.Split(contentStr, "\n")

	wordSearch := [][]string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineChars := strings.Split(line, "")
		wordSearch = append(wordSearch, lineChars)
	}

	return wordSearch, nil
}

func findXmas(search [][]string) int {
	maxy := len(search) - 1
	maxx := len(search[0]) - 1

	occurances := 0
	for r := range search {
		for c, cval := range search[r] {
			if cval == "X" {
				count := circularSearch(c, r, maxx, maxy, search)
				occurances += count
			}
		}
	}

	return occurances
}

func findCrossmas(search [][]string) int {
	maxy := len(search) - 1
	maxx := len(search[0]) - 1

	occurances := 0
	for r := range search {
		for c, cval := range search[r] {
			if cval == "A" {
				if checkCross(c, r, maxx, maxy, search) {
					occurances++
				}
			}
		}
	}

	return occurances
}

func circularSearch(x, y, maxx, maxy int, search [][]string) int {
	foundInstances := 0
	if checkUp(x, y, maxx, maxy, search) {
		// fmt.Printf("found upward at %d,%d\n", x, y)
		foundInstances++
	}

	if checkRightDiagUp(x, y, maxx, maxy, search) {
		// fmt.Printf("found upward right diagonal at %d,%d\n", x, y)
		foundInstances++
	}

	if checkRight(x, y, maxx, maxy, search) {
		// fmt.Printf("found right at %d,%d\n", x, y)
		foundInstances++
	}

	if checkRightDiagDown(x, y, maxx, maxy, search) {
		// fmt.Printf("found downward right diagonal at %d,%d\n", x, y)
		foundInstances++
	}

	if checkDown(x, y, maxx, maxy, search) {
		// fmt.Printf("found down at %d,%d\n", x, y)
		foundInstances++
	}

	if checkLeftDiagDown(x, y, maxx, maxy, search) {
		// fmt.Printf("found downward left diagonal at %d,%d\n", x, y)
		foundInstances++
	}

	if checkLeft(x, y, maxx, maxy, search) {
		// fmt.Printf("found left at %d,%d\n", x, y)
		foundInstances++
	}

	if checkLeftDiagUp(x, y, maxx, maxy, search) {
		// fmt.Printf("found upward left diagonal at %d,%d\n", x, y)
		foundInstances++
	}

	return foundInstances
}

func checkUp(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idx := y - (i + 1)
		if idx < 0 {
			return false
		}
		if search[idx][x] != q {
			return false
		}
	}

	return true
}

func checkDown(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idx := y + (i + 1)
		if idx > maxy {
			return false
		}
		if search[idx][x] != q {
			return false
		}
	}

	return true
}

func checkRight(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idx := x + (i + 1)
		if idx > maxx {
			return false
		}
		if search[y][idx] != q {
			return false
		}
	}

	return true
}

func checkLeft(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idx := x - (i + 1)
		if idx < 0 {
			return false
		}
		if search[y][idx] != q {
			return false
		}
	}

	return true
}

func checkRightDiagUp(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idxy := y - (i + 1)
		idxx := x + (i + 1)
		if idxy < 0 || idxx > maxx {
			return false
		}
		if search[idxy][idxx] != q {
			return false
		}
	}

	return true
}

func checkRightDiagDown(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idxy := y + (i + 1)
		idxx := x + (i + 1)
		if idxy > maxy || idxx > maxx {
			return false
		}
		if search[idxy][idxx] != q {
			return false
		}
	}

	return true
}

func checkLeftDiagUp(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idxy := y - (i + 1)
		idxx := x - (i + 1)
		if idxy < 0 || idxx < 0 {
			return false
		}
		if search[idxy][idxx] != q {
			return false
		}
	}

	return true
}

func checkLeftDiagDown(x, y, maxx, maxy int, search [][]string) bool {
	for i, q := range SEARCHES {
		idxy := y + (i + 1)
		idxx := x - (i + 1)
		if idxy > maxy || idxx < 0 {
			return false
		}
		if search[idxy][idxx] != q {
			return false
		}
	}

	return true
}

func checkCross(x, y, maxx, maxy int, search [][]string) bool {
	tlIdy, tlIdx := y-1, x-1
	blIdy, blIdx := y+1, x-1
	trIdy, trIdx := y-1, x+1
	brIdy, brIdx := y+1, x+1

	if (tlIdx < 0 || tlIdy < 0) || (blIdx < 0 || blIdy > maxy) || (trIdx > maxx || trIdy < 0) || (brIdx > maxx || brIdy > maxy) {
		return false
	}

	tl := search[tlIdy][tlIdx]
	bl := search[blIdy][blIdx]
	tr := search[trIdy][trIdx]
	br := search[brIdy][brIdx]

	return checkDiag(tl, br) && checkDiag(bl, tr)

}

func checkDiag(l, r string) bool {
	return (l == "M" && r == "S") || (l == "S" && r == "M")
}
