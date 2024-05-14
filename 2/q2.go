package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToString(resInt []int) string {
	var resStr []string
	for _, num := range resInt {
		strInt := strconv.Itoa(num)
		resStr = append(resStr, strInt)
	}
	resFinal := strings.Join(resStr, "")
	return resFinal
}

func decode(input string) (string, error) {
	splitInput := strings.Split(input, "")
	var res []int
	bp := 0

	res = append(res, 0)

	for i := range input {

		if splitInput[i] == "L" {
			if res[i] == 0 {
				for j := len(res) - 1; j >= bp; j-- {
					res[j] += 1
				}
			}
			res = append(res, 0)
		}

		if splitInput[i] == "R" {
			bp = i + 1
			next := res[i] + 1
			res = append(res, next)
		}

		if splitInput[i] == "=" {
			current := res[i]
			res = append(res, current)
		}
	}

	result := convertToString(res)
	return result, nil
}

func main() {
	var input string

    fmt.Print("Input: ")

    _, err := fmt.Scanln(&input)

    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    
	res, err := decode(input)

	fmt.Println("Output: ", res)
}