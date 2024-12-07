package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type direction int

type Guard struct {
	Direction direction
	POSx      int
	POSy      int
	Visited   [][2]int
	VisitMap  map[string]int
}

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
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
	labMap := getInput()

	// for _, line := range labMap {
	// 	fmt.Printf("%+v\n", line)
	// }

	x, y := findStart(labMap)

	if x == -1 || y == -1 {
		panic("could not find starting point")
	}
	guard := &Guard{
		Direction: UP,
		POSx:      x,
		POSy:      y,
		Visited:   [][2]int{},
		VisitMap:  map[string]int{},
	}

	RunRoute(guard, labMap)

	for _, pos := range guard.Visited {
		fmt.Printf("%+v\n", pos)
	}

	fmt.Println(len(guard.Visited))
}

func part2() {
	labMap := getInput()

	// for _, line := range labMap {
	// 	fmt.Printf("%+v\n", line)
	// }

	x, y := findStart(labMap)

	if x == -1 || y == -1 {
		panic("could not find starting point")
	}
	guard := &Guard{
		Direction: UP,
		POSx:      x,
		POSy:      y,
		Visited:   [][2]int{},
		VisitMap:  map[string]int{},
	}

	RunRoute(guard, labMap)

	blockerOptions := guard.Visited[1:]

	count := CheckPositions(labMap, blockerOptions, x, y)
	fmt.Println(count)
}

func getInput() [][]string {
	content, err := os.ReadFile("input.txt")
	// content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	labMap := [][]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		positions := strings.Split(line, "")
		labMap = append(labMap, positions)
	}

	return labMap
}

// GUARD funcs [

func (g *Guard) IsOffMap(maxx, maxy int) bool {
	if g.POSx > maxx || g.POSx < 0 {
		return true
	}

	if g.POSy > maxy || g.POSy < 0 {
		return true
	}

	return false
}

func (g *Guard) Move() {
	pos := [2]int{g.POSx, g.POSy}
	g.Visited = uniqueAppend(g.Visited, pos)
	g.VisitMap[fmt.Sprintf("%d|%d", g.POSx, g.POSy)]++
	switch g.Direction {
	case UP:
		g.POSy -= 1
	case RIGHT:
		g.POSx += 1
	case DOWN:
		g.POSy += 1
	case LEFT:
		g.POSx -= 1
	default:
		fmt.Println("unknown direction")
	}

}

func (g *Guard) Rotate() {
	g.Direction = (g.Direction + 1) % 4
}

func (g *Guard) Log() {
	fmt.Printf("Pos: %d,%d Direction: %d Traveled: %d\n", g.POSx, g.POSy, g.Direction, len(g.Visited))
}

// ] end guard funcs

func findStart(labMap [][]string) (int, int) {
	for r := range labMap {
		for c, pos := range labMap[r] {
			if pos == "^" {
				return c, r
			}
		}
	}

	return -1, -1
}

func RunRoute(guard *Guard, labMap [][]string) error {
	maxx, maxy := len(labMap[0])-1, len(labMap)-1
	for !guard.IsOffMap(maxx, maxy) {
		// guard.Log()
		if !CheckNext(guard, labMap) {
			guard.Rotate()
			continue
		}
		guard.Move()
		if checkForLoop(guard) {
			return errors.New("hit loop")
		}
	}

	return nil
}

func CheckNext(guard *Guard, labMap [][]string) bool {
	x, y := guard.POSx, guard.POSy

	switch guard.Direction {
	case UP:
		y -= 1
	case RIGHT:
		x += 1
	case DOWN:
		y += 1
	case LEFT:
		x -= 1
	}

	if x > len(labMap[0])-1 || x < 0 {
		return true
	}

	if y > len(labMap)-1 || y < 0 {
		return true
	}

	next := labMap[y][x]

	switch next {
	case "#":
		return false
	default:
		return true
	}
}

func uniqueAppend(list [][2]int, pos [2]int) [][2]int {
	for _, item := range list {
		if pos[0] == item[0] && pos[1] == item[1] {
			return list
		}
	}

	return append(list, pos)
}

func CheckPositions(labMap [][]string, positions [][2]int, startX, startY int) int {
	count := 0
	wg := sync.WaitGroup{}
	finishedCount := 0
	for i, pos := range positions {
		// fmt.Printf("Trying position %d,%d\n", pos[0], pos[1])
		wg.Add(1)
		go func() {
			fmt.Printf("starting position %d\n", i)
			tmpMap := copyMap(labMap)
			tmpMap[pos[1]][pos[0]] = "#"

			// for r := range labMap {
			// 	fmt.Printf("%+v %+v\n", tmpMap[r], labMap[r])
			// }

			guard := &Guard{
				Direction: UP,
				POSx:      startX,
				POSy:      startY,
				Visited:   [][2]int{},
				VisitMap:  map[string]int{},
			}

			err := RunRoute(guard, tmpMap)
			if err != nil {
				count++
			}
			wg.Done()
			finishedCount++
			fmt.Printf("finished position %d total: %d count: %d\n", i, finishedCount, count)
		}()
	}

	wg.Wait()

	return count
}

func checkForLoop(guard *Guard) bool {
	count := 0
	for _, v := range guard.VisitMap {
		if v > 100 {
			count++
		}
	}

	if count > 0 {
		return true
	}

	return false
}

func copyMap(inMap [][]string) [][]string {
	out := [][]string{}

	for r := range inMap {
		tmpRow := []string{}
		tmpRow = append(tmpRow, inMap[r]...)
		out = append(out, tmpRow)
	}

	return out
}
