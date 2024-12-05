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
var movesLeftRight = [][]int{
	[]int{-1, -1},
	[]int{1, 1},
}
var movesRightLeft = [][]int{
	[]int{-1, 1},
	[]int{1, -1},
}

var total = 0

func Run_2(input string) int {
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
			if matrix[i][j] == "A" {
				fmt.Printf("found A on: %v%v\n", i, j)

				found := map[string]int{"M": 0, "S": 0}
				// var coords string = ""
				for _, m := range movesLeftRight {
					row := i + m[0]
					col := j + m[1]

					if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 {
						if _, ok := found[matrix[row][col]]; ok {
							// fmt.Println("found letter:", matrix[row][col])
							// fmt.Println("count letter:", found[matrix[row][col]])
							// coords += fmt.Sprintf("%v%v ", row, col)
							found[matrix[row][col]]++
						}
						// fmt.Printf("founx M on: %v%v\n", row, col)
						// Search_2(row, col, 1, fmt.Sprintf("%v%v %v%v", i, j, row, col))
					}
				}

				if found["M"] != 1 && found["S"] != 1 {
					continue
				}

				for _, m := range movesRightLeft {
					row := i + m[0]
					col := j + m[1]

					if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 {
						if _, ok := found[matrix[row][col]]; ok {
							// fmt.Println("found letter:", matrix[row][col])
							// fmt.Println("count letter:", found[matrix[row][col]])
							// coords += fmt.Sprintf("%v%v ", row, col)
							found[matrix[row][col]]++
						}
						// fmt.Printf("founx M on: %v%v\n", row, col)
						// Search_2(row, col, 1, fmt.Sprintf("%v%v %v%v", i, j, row, col))
					}
				}

				if found["M"] == 2 && found["S"] == 2 {
					// fmt.Println("found", found)
					// fmt.Println("coords", coords)
					total++
				}

			}
		}
	}

	return total
}

func Run_1(input string) int {
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
						Search_1(row, col, rowVec, colVec, 2, fmt.Sprintf("%v%v %v%v", i, j, row, col))
					}
				}
			}
		}
	}

	return total
}

func Search_1(row, col, rowVec, colVec, xmasIdx int, coords string) {
	if xmasIdx == len(word) {
		// fmt.Println("coords", coords)
		total++
		return
	}
	row = row + rowVec
	col = col + colVec
	if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 && matrix[row][col] == word[xmasIdx] {
		Search_1(row, col, rowVec, colVec, xmasIdx+1, coords+" "+fmt.Sprintf("%v%v", row, col))
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
