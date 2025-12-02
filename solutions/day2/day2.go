package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func main() {

	data, err := readInput("solutions\\day2\\day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ans1, ans2 := 0, 0
	elapsed1, elapsed2 := time.Duration(0), time.Duration(0)

	for _, part := range strings.Split(data, ",") {
		rangeParts := strings.Split(part, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			n := len(s)

			// Part 1
			start1 := time.Now()

			if n%2 == 0 {
				mid := n / 2
				if s[:mid] == s[mid:] {
					ans1 += i
				}
			}

			elapsed1 = time.Since(start1)

			// Part 2
			start2 := time.Now()

			for l := 1; l <= n/2; l++ {
				if n%l != 0 {
					continue
				}

				ok := true

				for j := l; j < n; j++ {
					if s[j] != s[j%l] {
						ok = false
						break
					}
				}

				if ok {
					ans2 += i
					break
				}
			}
			elapsed2 = time.Since(start2)

		}
	}

	fmt.Printf("Part 1: %d (%d ms)\n", ans1, elapsed1.Milliseconds())
	fmt.Printf("Part 2: %d (%d ms)\n", ans2, elapsed2.Milliseconds())
}
