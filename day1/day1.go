package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {
	f, err := os.Open("day1/data.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	l, r := []int{}, []int{}
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "   ")
		fd, err := strconv.Atoi(res[0])

		if err != nil {
			panic(err)
		}
		sd, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		l = append(l, fd)
		r = append(r, sd)
	}

	change := true
	for change == true {
		change = false
		for i := range l {
			if i == 0 {
				continue
			}
			if l[i] < l[i-1] {
				l[i], l[i-1] = l[i-1], l[i]
				change = true
			}
		}
	}

	change = true
	for change == true {
		change = false
		for i := range r {
			if i == 0 {
				continue
			}
			if r[i] < r[i-1] {
				r[i], r[i-1] = r[i-1], r[i]
				change = true
			}
		}
	}

	total := 0
	for i := range l {
		total += abs(l[i], r[i])
	}

	fmt.Println(total)
}

func abs(f, s int) int {
	if f > s {
		return f - s
	}
	return s - f
}
