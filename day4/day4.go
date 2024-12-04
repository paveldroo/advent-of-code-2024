package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var word = []string{"X", "M", "A", "S"}

var moves = [][]int{
	[]int{-1, 0},
	[]int{1, 0},
	[]int{0, -1},
	[]int{0, 1},
	[]int{-1, -1},
	[]int{-1, 1},
	[]int{1, -1},
	[]int{1, 1},
}

var matrix = [][]string{}

var total = 0

func Run(input string) int {
	if input == "" {
		input = Parse()
	}

	for _, m := range strings.Split(input, "\n") {
		if len(m) == 0 {
			continue
		}
		matrix = append(matrix, strings.Split(m, ""))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "X" {
				// fmt.Printf("founx X on: %v%v\n", i, j)
				for _, m := range moves {
					rowVec := m[0]
					colVec := m[1]
					row := i + rowVec
					col := j + colVec
					if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 && matrix[row][col] == "M" {
						// fmt.Printf("founx M on: %v%v\n", row, col)
						Search(row, col, rowVec, colVec, 2, fmt.Sprintf("%v%v %v%v", i, j, row, col))
					}
				}
			}
		}
	}

	return total
}

func Search(row, col, rowVec, colVec, xmasIdx int, coords string) {
	if xmasIdx == len(word) {
		// fmt.Println("coords", coords)
		total++
		return
	}
	row = row + rowVec
	col = col + colVec
	if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 && matrix[row][col] == word[xmasIdx] {
		Search(row, col, rowVec, colVec, xmasIdx+1, coords+" "+fmt.Sprintf("%v%v", row, col))
	}
}

func Parse() string {
	f, err := os.Open("day4/data.txt")
	if err != nil {
		panic(err)
	}

	res := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		res = res + row + "\n"
	}
	return res
}
