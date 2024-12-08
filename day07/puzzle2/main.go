package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maja42/goval"
	"github.com/suchen-sci/gfn"
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

	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		eq := strings.Split(line, ":")
		testResult := eq[0]
		parts := strings.Fields(strings.TrimSpace(eq[1]))
		for i := 0; i < intPow(3, (len(parts)-1)); i++ {
			signs := convertToTernary(i)
			signs = strings.ReplaceAll(signs, "0", "+")
			signs = strings.ReplaceAll(signs, "1", "*")
			signs = strings.ReplaceAll(signs, "2", "|")

			firstIteration := true
			exprResult := gfn.Reduce(parts, "1", func(acc, next string) string {
				var sign string
				if firstIteration {
					sign = "*"
				} else {
					sign = pop(&signs)
				}

				firstIteration = false

				if sign == "|" {
					return acc + next
				}

				eval := goval.NewEvaluator()
				result, err := eval.Evaluate(acc+sign+next, nil, nil)
				if err != nil {
					panic(err)
				}

				return strconv.Itoa(result.(int))
			})
			if strToInt(exprResult) == strToInt(testResult) {
				sum += strToInt(testResult)
				break
			}
		}
	}
	fmt.Println(sum)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func pop(alist *string) string {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return string(rv)
}

func convertToTernary(n int) string {
	if n == 0 {
		return "000000000000000000000000000000000"
	}
	ternary := ""
	for n > 0 {
		ternary = strconv.Itoa(n%3) + ternary
		n /= 3
	}
	return "000000000000000000000000000000000" + ternary
}
