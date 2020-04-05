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
	lines := mustScanInt(scanner)
	for i := 0; i < lines; i++ {
		fmt.Fprintf(out, "Case #%v: %v\n", i+1, parenthify(mustScanIntArray(scanner)))
	}
}

func parenthify(arr []int) string {
	var result string
	for i, num := range arr {
		var prev int
		if i > 0 {
			prev = arr[i-1]
		}
		var next int
		if i < len(arr)-1 {
			next = arr[i+1]
		}
		if num > prev {
			for i := 0; i < num-prev; i++ {
				result += "("
			}
		}
		result += strconv.Itoa(num)
		if num > next {
			for i := 0; i < num-next; i++ {
				result += ")"
			}
		}
	}
	return result
}

func mustScanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func mustScanIntArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	var result []int
	for _, rawInt := range strings.Split(scanner.Text(), "") {
		num, _ := strconv.Atoi(rawInt)
		result = append(result, num)
	}
	return result
}
