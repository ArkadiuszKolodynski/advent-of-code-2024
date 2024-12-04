package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode/utf8"
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

	pattern := regexp.MustCompile("XMAS")
	rows := make([]string, 0)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)

		// find pattern occurrences horizontally
		count += len(pattern.FindAllString(line, -1))
		count += len(pattern.FindAllString(reverse(line), -1))
	}

	for i := 0; i < len(rows[0]); i++ {
		col := ""
		for j := 0; j < len(rows); j++ {
			col += string(rows[j][i])
		}
		// find pattern occurrences vertically
		count += len(pattern.FindAllString(col, -1))
		count += len(pattern.FindAllString(reverse(col), -1))
	}

	for i := 0; i < len(rows); i++ {
		leading_diagonal_upper := ""
		leading_diagonal_lower := ""
		antidiagonal_upper := ""
		antidiagonal_lower := ""
		for j := 0; j < len(rows)-i; j++ {
			leading_diagonal_upper += string(rows[j][j+i])
			leading_diagonal_lower += string(rows[j+i][j])
			antidiagonal_upper += string(rows[j][len(rows)-j-i-1])
			antidiagonal_lower += string(rows[j+i][len(rows)-j-1])
		}

		// find pattern occurrences diagonally
		if len(leading_diagonal_upper) >= 3 {
			count += len(pattern.FindAllString(leading_diagonal_upper, -1))
			count += len(pattern.FindAllString(reverse(leading_diagonal_upper), -1))
		}

		if i > 0 && len(leading_diagonal_lower) >= 3 {
			count += len(pattern.FindAllString(leading_diagonal_lower, -1))
			count += len(pattern.FindAllString(reverse(leading_diagonal_lower), -1))
		}

		if len(antidiagonal_upper) >= 3 {
			count += len(pattern.FindAllString(antidiagonal_upper, -1))
			count += len(pattern.FindAllString(reverse(antidiagonal_upper), -1))
		}

		if i > 0 && len(antidiagonal_lower) >= 3 {
			count += len(pattern.FindAllString(antidiagonal_lower, -1))
			count += len(pattern.FindAllString(reverse(antidiagonal_lower), -1))
		}
	}

	fmt.Println(count)
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
