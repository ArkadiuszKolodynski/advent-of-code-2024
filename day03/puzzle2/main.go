package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	buffer, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	memory := string(buffer)
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	markerRe := regexp.MustCompile(`(?:do\(\)|don't\(\))`)
	mulMatches := mulRe.FindAllStringIndex(memory, -1)
	result := 0

	for _, mulMatch := range mulMatches {
		mulStart := mulMatch[0]
		mulStr := memory[mulMatch[0]:mulMatch[1]]

		markerMatches := markerRe.FindAllStringIndex(memory[:mulStart], -1)

		var lastMarker string
		if len(markerMatches) > 0 {
			lastMarkerIdx := markerMatches[len(markerMatches)-1]
			lastMarker = memory[lastMarkerIdx[0]:lastMarkerIdx[1]]
		} else {
			lastMarker = ``
		}

		if lastMarker != `don't()` {
			tuple := strings.Split(mulStr[4:len(mulStr)-1], ",")
			result += stringToInt(tuple[0]) * stringToInt(tuple[1])
		}
	}

	fmt.Println(result)
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
