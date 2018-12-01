package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func urlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func fileToLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func main() {
	lines, _ := fileToLines("input.txt")
	var m = make(map[int]bool)
	i1 := 0
	i2 := 0
	found := false
	idx := 0
	i := 0
	for !found {
		for _, line := range lines {
			x, _ := strconv.Atoi(line)
			i += x
			if _, ok := m[i]; ok && !found {
				i2 = i
				found = true
			}
			m[i] = true
		}
		idx++
		if idx == 1 {
			i1 = i
		}
	}

	fmt.Printf("Day 1 Part 1 %v, Part 2 %v", i1, i2)
}
