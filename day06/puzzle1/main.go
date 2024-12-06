package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	guardIndicator := rune('^')
	blockerIndicator := rune('#')
	directions := []string{"up", "right", "down", "left"}
	currentDirection := 0
	var currentPosition []int
	grid := make([][]rune, 0)
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
		if strings.Contains(line, string(guardIndicator)) {
			currentPosition = []int{i, strings.Index(line, string(guardIndicator))}
		}
		i++
	}

	steps := make(map[string]bool)
	steps[fmt.Sprintf("%d,%d", currentPosition[0], currentPosition[1])] = true
	for {
		nextPosition := make([]int, 2)
		switch directions[currentDirection] {
		case "up":
			nextPosition = []int{currentPosition[0] - 1, currentPosition[1]}
		case "right":
			nextPosition = []int{currentPosition[0], currentPosition[1] + 1}
		case "down":
			nextPosition = []int{currentPosition[0] + 1, currentPosition[1]}
		case "left":
			nextPosition = []int{currentPosition[0], currentPosition[1] - 1}
		}

		if nextPosition[0] < 0 || nextPosition[0] >= len(grid) || nextPosition[1] < 0 || nextPosition[1] >= len(grid[0]) {
			break
		}

		if grid[nextPosition[0]][nextPosition[1]] == blockerIndicator {
			currentDirection = (currentDirection + 1) % 4
			continue
		}

		currentPosition = nextPosition
		steps[fmt.Sprintf("%d,%d", currentPosition[0], currentPosition[1])] = true
	}

	fmt.Println(len(steps))
}
