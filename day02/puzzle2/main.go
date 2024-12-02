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

		isSafe := isReportSafe(report)

		if !isSafe {
			for i := 0; i < len(report); i++ {
				tmpReport := make([]int, len(report))
				copy(tmpReport, report)
				newReport := remove(tmpReport, i)
				if isReportSafe(newReport) {
					isSafe = true
					break
				}
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

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}
	isAscending := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		if (isAscending && report[i-1] >= report[i]) ||
			(!isAscending && report[i-1] <= report[i]) ||
			max(report[i]-report[i-1], -(report[i]-report[i-1])) > 3 {
			return false
		}
	}
	return true
}

func appendNumberToList(id string, list []int) []int {
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	list = append(list, n)
	return list
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
