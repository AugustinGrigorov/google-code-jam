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

func main() {
	scanner := bufio.NewScanner(in)
	testCases := mustScanInt(scanner)
	for i := 0; i < testCases; i++ {
		numberOfActivites := mustScanInt(scanner)
		activities := make([][2]int, numberOfActivites)
		for j := 0; j < numberOfActivites; j++ {
			activities[j] = mustScanActivity(scanner)
		}
		fmt.Fprintf(out, "Case #%v: %v\n", i+1, computeSchedule(activities))
	}
}

func computeSchedule(activities [][2]int) string {
	schedule := make([]string, len(activities))
	availability := map[string]int{
		"C": 0,
		"J": 0,
	}
	for i := 0; i <= 1440; i++ {
		for activityIndex, activity := range activities {
			if activity[1] == i {
				if availability["C"] == activityIndex+1 {
					availability["C"] = 0
				}
				if availability["J"] == activityIndex+1 {
					availability["J"] = 0
				}
			}
		}
		for activityIndex, activity := range activities {
			if activity[0] <= i && activity[1] > i {
				if !(availability["C"] == activityIndex+1 || availability["J"] == activityIndex+1) {
					if availability["C"] == 0 {
						availability["C"] = activityIndex + 1
						schedule[activityIndex] = "C"
					} else if availability["J"] == 0 {
						availability["J"] = activityIndex + 1
						schedule[activityIndex] = "J"
					} else {
						return "IMPOSSIBLE"
					}
				}
			}
		}
	}
	return strings.Join(schedule, "")
}

func mustScanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func mustScanActivity(scanner *bufio.Scanner) [2]int {
	scanner.Scan()
	var result [2]int
	for i, rawInt := range strings.Split(scanner.Text(), " ") {
		num, _ := strconv.Atoi(rawInt)
		result[i] = num
	}
	return result
}
