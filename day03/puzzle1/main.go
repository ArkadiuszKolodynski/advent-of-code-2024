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
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	muls := re.FindAllString(memory, -1)

	result := 0
	for _, mul := range muls {
		tuple := strings.Split(mul[4:len(mul)-1], ",")
		result += stringToInt(tuple[0]) * stringToInt(tuple[1])
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
