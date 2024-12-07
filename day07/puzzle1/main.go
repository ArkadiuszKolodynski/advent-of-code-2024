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
		for i := 0; i < intPow(2, (len(parts)-1)); i++ {
			signs := fmt.Sprintf("%0"+strconv.Itoa(len(parts)-1)+"b", i)
			signs = strings.ReplaceAll(signs, "0", "+")
			signs = strings.ReplaceAll(signs, "1", "*")

			firstIteration := true
			exprResult := gfn.Reduce(parts, "1", func(acc, next string) string {
				var sign string
				if firstIteration {
					sign = "*"
				} else {
					sign = pop(&signs)
				}

				eval := goval.NewEvaluator()
				result, err := eval.Evaluate(acc+string(sign)+next, nil, nil)
				if err != nil {
					panic(err)
				}

				firstIteration = false
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
