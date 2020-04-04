package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix [][]int

func main() {
	// Parse input
	scanner := bufio.NewScanner(os.Stdin)
	testCases := mustScanInt(scanner)
	var matrices []matrix
	for i := 0; i < testCases; i++ {
		matrixSize := mustScanInt(scanner)
		var newMatrix matrix
		for m := 0; m < matrixSize; m++ {
			row := mustScanIntArray(scanner)
			newMatrix = append(newMatrix, row)
		}
		matrices = append(matrices, newMatrix)
	}

	// Compute matrix stats
	for testCase, m := range matrices {
		var ms = len(m)
		var trace int
		var rowsWithRepeatedElements int
		var columnsWithRepeatedElements int
		for i := 0; i < ms; i++ {
			trace += m[i][i]
		}
		for i := 0; i < ms; i++ {
			rowNumberCounts := make(map[int]int, ms)
			colNumberCounts := make(map[int]int, ms)
			for j := 0; j < ms; j++ {
				rowNumberCounts[m[i][j]]++
				colNumberCounts[m[j][i]]++
			}
			for _, occurances := range rowNumberCounts {
				if occurances > 1 {
					rowsWithRepeatedElements++
					break
				}
			}
			for _, occurances := range colNumberCounts {
				if occurances > 1 {
					columnsWithRepeatedElements++
					break
				}
			}
		}
		fmt.Printf("Case #%v: %v %v %v\n", testCase+1, trace, rowsWithRepeatedElements, columnsWithRepeatedElements)
	}
}

func mustScanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func mustScanIntArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	var result []int
	for _, rawInt := range strings.Split(scanner.Text(), " ") {
		num, _ := strconv.Atoi(rawInt)
		result = append(result, num)
	}
	return result
}
