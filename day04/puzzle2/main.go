package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	rows := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}

	count := 0
	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i])-1; j++ {
			if rows[i][j] == 'A' {
				topLeft := string(rows[i-1][j-1])
				topRight := string(rows[i-1][j+1])
				bottomLeft := string(rows[i+1][j-1])
				bottomRight := string(rows[i+1][j+1])

				dict := make(map[string]int)
				dict[topLeft] = dict[topLeft] + 1
				dict[topRight] = dict[topRight] + 1
				dict[bottomLeft] = dict[bottomLeft] + 1
				dict[bottomRight] = dict[bottomRight] + 1

				if topLeft != bottomRight && topRight != bottomLeft && dict["M"] == 2 && dict["S"] == 2 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
