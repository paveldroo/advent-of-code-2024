package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var secondChanseInc = [][]int{}
var secondChanseDec = [][]int{}

func First() {
	f, err := os.Open("day2/data.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	incSlice := [][]int{}
	decSlice := [][]int{}

	for s.Scan() {
		a := strings.Split(s.Text(), " ")
		ints := makeInts(a)
		if isInc(ints) {
			incSlice = append(incSlice, ints)
		} else {
			decSlice = append(decSlice, ints)
		}
	}

	total := 0

	for _, s := range incSlice {
		if ValidateInc(s, false) {
			total++
		}
	}

	for _, s := range decSlice {
		if ValidateDec(s, false) {
			total++
		}
	}

	fmt.Println(total)
}

func makeInts(d []string) []int {
	var res []int
	for i := range d {
		v, err := strconv.Atoi(d[i])
		if err != nil {
			panic(err)
		}
		res = append(res, v)
	}
	return res
}

func isInc(d []int) bool {
	inc := 0
	dec := 0
	for i := range d {
		if i == 0 {
			continue
		}

		if d[i] > d[i-1] {
			inc++
		} else if d[i] < d[i-1] {
			dec++
		}
	}

	if inc > dec {
		return true
	}

	return false
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func ValidateInc(d []int, sec bool) bool {
	dc := make([]int, len(d))
	copy(dc, d)
	for i := range dc {
		if i < len(dc)-1 {
			cur := dc[i]
			next := dc[i+1]
			if next <= cur {
				if !sec {
					ss1 := make([]int, len(dc))
					copy(ss1, dc)
					s1 := remove(ss1, i)
					fmt.Println("s1", s1)

					ss2 := make([]int, len(dc))
					copy(ss2, dc)
					s2 := remove(ss2, i+1)
					fmt.Println("s2", s2)
					return ValidateInc(s1, true) || ValidateInc(s2, true)
				}
				return false
			}

			if next-cur > 3 {
				if !sec {
					ss1 := make([]int, len(dc))
					copy(ss1, dc)
					s1 := remove(ss1, i)
					fmt.Println("s1", s1)

					ss2 := make([]int, len(dc))
					copy(ss2, dc)
					s2 := remove(ss2, i+1)
					fmt.Println("s2", s2)
					return ValidateInc(s1, true) || ValidateInc(s2, true)
				}
				return false
			}
		}
	}
	return true
}

func ValidateDec(d []int, sec bool) bool {
	dc := make([]int, len(d))
	copy(dc, d)
	for i := range dc {
		if i < len(dc)-1 {
			cur := dc[i]
			next := dc[i+1]
			if next >= cur {
				if !sec {
					ss1 := make([]int, len(dc))
					copy(ss1, dc)
					s1 := remove(ss1, i)
					fmt.Println("s1", s1)

					ss2 := make([]int, len(dc))
					copy(ss2, dc)
					s2 := remove(ss2, i+1)
					fmt.Println("s2", s2)
					return ValidateDec(s1, true) || ValidateDec(s2, true)
				}
				return false
			}

			if next-cur < -3 {
				if !sec {
					ss1 := make([]int, len(dc))
					copy(ss1, dc)
					s1 := remove(ss1, i)
					fmt.Println("s1", s1)

					ss2 := make([]int, len(dc))
					copy(ss2, dc)
					s2 := remove(ss2, i+1)
					fmt.Println("s2", s2)
					return ValidateDec(s1, true) || ValidateDec(s2, true)
				}
				return false
			}
		}
	}
	return true
}
