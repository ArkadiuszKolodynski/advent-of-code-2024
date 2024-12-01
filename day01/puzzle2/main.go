package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	leftList := make([]int, 0)
	rightList := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		leftList = appendNumberFromID(ids[0], leftList)
		rightList = appendNumberFromID(ids[1], rightList)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	similarity := 0
	for _, n := range leftList {
		s := 0
		for _, m := range rightList {
			if n == m {
				s++
			}
		}
		similarity += n * s
	}
	fmt.Println(similarity)
}

func appendNumberFromID(id string, list []int) []int {
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	list = append(list, n)
	return list
}
