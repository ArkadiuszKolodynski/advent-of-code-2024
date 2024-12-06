package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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

	rules := make(map[string][]string)
	var updates []string
	isFirstBlock := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isFirstBlock = false
			continue
		}

		if isFirstBlock {
			rulePart := strings.Split(line, "|")
			rules[rulePart[0]] = append(rules[rulePart[0]], rulePart[1])
		} else {
			updates = append(updates, line)
		}
	}

	sum := 0
	for _, update := range updates {
		updateIsValid := true
		splittedUpdate := strings.Split(update, ",")
		for i := len(splittedUpdate) - 1; i >= 1; i-- {
			if rules[splittedUpdate[i]] == nil {
				continue
			}
			for j := 0; j < i; j++ {
				if slices.Contains((rules[splittedUpdate[i]]), splittedUpdate[j]) {
					updateIsValid = false
					break
				}
			}
			if !updateIsValid {
				break
			}
		}
		if updateIsValid {
			sum += stringToInt(splittedUpdate[len(splittedUpdate)/2])
		}
	}

	fmt.Println(sum)
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
