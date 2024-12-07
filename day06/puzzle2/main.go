package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	guardIndicator := rune('^')
	blockerIndicator := rune('#')
	directions := []string{"up", "right", "down", "left"}
	startingDirection := 0
	var startingPosition []int
	grid := make([][]rune, 0)
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
		if strings.Contains(line, string(guardIndicator)) {
			startingPosition = []int{i, strings.Index(line, string(guardIndicator))}
		}
		i++
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == blockerIndicator || grid[i][j] == guardIndicator {
				continue
			}

			gridCopy := make([][]rune, len(grid))
			for i := range grid {
				gridCopy[i] = make([]rune, len(grid[i]))
				copy(gridCopy[i], grid[i])
			}
			gridCopy[i][j] = blockerIndicator

			currentDirection := startingDirection
			currentPosition := startingPosition

			steps := make(map[string]bool)
			steps[fmt.Sprintf("%d,%d,%s", currentPosition[0], currentPosition[1], directions[currentDirection])] = true
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

				if nextPosition[0] < 0 || nextPosition[0] >= len(gridCopy) || nextPosition[1] < 0 || nextPosition[1] >= len(gridCopy[0]) {
					break
				}

				if gridCopy[nextPosition[0]][nextPosition[1]] == blockerIndicator {
					currentDirection = (currentDirection + 1) % 4
					continue
				}

				currentPosition = nextPosition
				if steps[fmt.Sprintf("%d,%d,%s", currentPosition[0], currentPosition[1], directions[currentDirection])] {
					count++
					break
				}
				steps[fmt.Sprintf("%d,%d,%s", currentPosition[0], currentPosition[1], directions[currentDirection])] = true
			}
		}
	}

	fmt.Println(count)
}
