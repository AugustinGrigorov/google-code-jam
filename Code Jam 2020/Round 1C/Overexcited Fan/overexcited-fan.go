package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout
var in io.Reader = os.Stdin

type coordinates struct {
	x, y int
}

type setup struct {
	startingPosition coordinates
	path             []string
}

func main() {
	// Parse input
	scanner := bufio.NewScanner(in)
	testCases := mustScanInt(scanner)
	for i := 0; i < testCases; i++ {
		fmt.Fprintf(out, "Case #%v: %v\n", i+1, checkRoute(parseSetup(scanner)))
	}
}

func checkRoute(setup setup) string {
	currentPosition := setup.startingPosition
	for moves, direction := range setup.path {
		currentPosition = move(currentPosition, direction)
		if abs(currentPosition.x)+abs(currentPosition.y) <= moves+1 {
			return strconv.Itoa(moves + 1)
		}
	}
	return "IMPOSSIBLE"
}

func move(position coordinates, direction string) coordinates {
	switch direction {
	case "N":
		return coordinates{
			x: position.x,
			y: position.y + 1,
		}
	case "E":
		return coordinates{
			x: position.x + 1,
			y: position.y,
		}
	case "W":
		return coordinates{
			x: position.x - 1,
			y: position.y,
		}
	case "S":
		return coordinates{
			x: position.x,
			y: position.y - 1,
		}
	}
	return position
}

func mustScanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func parseSetup(scanner *bufio.Scanner) setup {
	scanner.Scan()
	rawData := strings.Split(scanner.Text(), " ")
	x, _ := strconv.Atoi(rawData[0])
	y, _ := strconv.Atoi(rawData[1])
	return setup{
		startingPosition: coordinates{
			x: x,
			y: y,
		},
		path: strings.Split(rawData[2], ""),
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
