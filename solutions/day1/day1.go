package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}

func main() {
	filename := "solutions\\day1\\day1.txt"

	lines, err := readLines(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Part 1
	start1 := time.Now()

	dial1 := 50
	ans1 := 0

	for _, line := range lines {
		dir := line[0]
		val, _ := strconv.Atoi(line[1:])

		if dir == 'R' {
			dial1 = (dial1 + val) % 100
		} else {
			dial1 = (dial1 - val) % 100
			if dial1 < 0 {
				dial1 += 100
			}
		}

		if dial1 == 0 {
			ans1++
		}
	}

	elapsed1 := time.Since(start1)

	// Part 2
	start2 := time.Now()

	dial2 := 50
	ans2 := 0

	for _, line := range lines {
		dir := line[0]
		val, _ := strconv.Atoi(line[1:])
		step := 1
		if dir != 'R' {
			step = -1
		}

		for i := 0; i < val; i++ {
			dial2 += step
			dial2 %= 100
			if dial2 < 0 {
				dial2 += 100
			}
			if dial2 == 0 {
				ans2++
			}
		}
	}

	elapsed2 := time.Since(start2)

	fmt.Printf("Part 1: %d (%d ms)\n", ans1, elapsed1.Milliseconds())
	fmt.Printf("Part 2: %d (%d ms)\n", ans2, elapsed2.Milliseconds())
}
