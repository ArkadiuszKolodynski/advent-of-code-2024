package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := make([]int, 0)
	for i := 0; i < len(leftList); i++ {
		d := rightList[i] - leftList[i]
		distances = append(distances, max(d, -d))
	}

	distancesSum := 0
	for _, d := range distances {
		distancesSum += d
	}
	fmt.Println(distancesSum)
}

func appendNumberFromID(id string, list []int) []int {
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	list = append(list, n)
	return list
}
