package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines, nil
}

func computeMaxTwoDigit(s string) int {
	bestLeft := -1
	best := 0

	for _, battery := range s {
		if battery < '0' || battery > '9' {
			continue
		}
		d := int(battery - '0')

		if bestLeft != -1 {
			v := bestLeft*10 + d
			if v > best {
				best = v
			}
		}

		if d > bestLeft {
			bestLeft = d
		}
	}

	return best
}

func computeMaxKSubsequence(s string, k int) string {
	digits := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		}
	}

	n := len(digits)
	if k >= n {
		return string(digits)
	}

	result := make([]byte, 0, k)
	start := 0

	for remaining := k; remaining > 0; remaining-- {
		end := n - remaining

		bestDigit := byte('0')
		bestIndex := start

		for i := start; i <= end; i++ {
			if digits[i] > bestDigit {
				bestDigit = digits[i]
				bestIndex = i

				if bestDigit == '9' {
					break
				}
			}
		}

		result = append(result, bestDigit)
		start = bestIndex + 1
	}

	return string(result)
}

func main() {

	data, err := readInput("solutions/day3/day3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	const K = 12

	ans1 := 0
	var ans2 uint64 = 0

	for _, line := range data {
		max2 := computeMaxTwoDigit(line)
		ans1 += max2

		maxK := computeMaxKSubsequence(line, K)

		var val uint64
		for _, c := range maxK {
			val = val*10 + uint64(c-'0')
		}
		ans2 += val

	}

	fmt.Printf("Part 1: %d\n", ans1)
	fmt.Printf("Part 2: %d\n", ans2)
}
