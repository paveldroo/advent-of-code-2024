package day4

import (
	"bufio"
	"os"
)

func Run(input string) int {
	if input == "" {
		input = Parse()
	}

	return 0
}

func Parse() string {
	f, err := os.Open("day4/data.txt")
	if err != nil {
		panic(err)
	}

	res := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res += scanner.Text()
	}

	return res
}
