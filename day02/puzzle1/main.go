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

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		report := make([]int, 0)
		for _, id := range ids {
			report = appendNumberToList(id, report)
		}

		isAscending := report[0] < report[1]
		isSafe := true
		for i := 1; i < len(report); i++ {
			if (isAscending && report[i-1] >= report[i]) || (!isAscending && report[i-1] <= report[i]) || max(report[i]-report[i-1], -(report[i]-report[i-1])) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeReports++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(safeReports)
}

func appendNumberToList(id string, list []int) []int {
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	list = append(list, n)
	return list
}
