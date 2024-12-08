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

	dict := make(map[string][][]int, 0)
	rowsCount := 0
	var colsLength int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		chars := strings.Split(line, "")
		for colsLength = 0; colsLength < len(chars); colsLength++ {
			if line[colsLength] != '.' {
				dict[string(line[colsLength])] = append(dict[string(line[colsLength])], []int{rowsCount, colsLength})
			}
		}
		rowsCount++
	}

	antinodes := make(map[string]bool)
	for _, v := range dict {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				idiff := v[j][0] - v[i][0]
				jdiff := v[j][1] - v[i][1]
				antinode1 := []int{v[i][0] - idiff, v[i][1] - jdiff}
				antinode2 := []int{v[j][0] + idiff, v[j][1] + jdiff}
				if antinode1[0] >= 0 && antinode1[0] < rowsCount && antinode1[1] >= 0 && antinode1[1] < colsLength {
					antinodes[fmt.Sprintf("%d,%d", antinode1[0], antinode1[1])] = true
				}
				if antinode2[0] >= 0 && antinode2[0] < rowsCount && antinode2[1] >= 0 && antinode2[1] < colsLength {
					antinodes[fmt.Sprintf("%d,%d", antinode2[0], antinode2[1])] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}
