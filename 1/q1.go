package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

    data, err := ioutil.ReadFile("easy.json")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    var matrix [][]int

    err = json.Unmarshal(data, &matrix)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Println("input:", matrix)

	res := maxTotal(matrix)

	print("output: ", res)

}

func maxTotal(input [][]int) int {
	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			if input[i+1][j] > input[i+1][j+1] {
				input[i][j] += input[i+1][j]
			} else {
				input[i][j] += input[i+1][j+1]	
			}
		}
	}
	return input[0][0]
}
