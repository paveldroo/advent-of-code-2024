package day3

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var p = []rune{'m', 'u', 'l', '('}
var d = []rune{'d', 'o', '(', ')'}
var dn = []rune{'d', 'o', 'n', 39, 't', '(', ')'}

func Parse() {
	f, err := os.Open("day3/data.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	res := 0
	for s.Scan() {
		r := First(s.Text())
		res += r
	}

	fmt.Println(res)
}

func First(input string) int {
	ri := []rune(input)
	do := true
	total := 0
	for i, v := range ri {
		if v == 'd' {
			ok := ProcessDo(append([]rune(nil), ri[i:]...))
			if ok {
				do = true
			}
			ok = ProcessDont(append([]rune(nil), ri[i:]...))
			if ok {
				do = false
			}
		}

		if v == 'm' && do {
			ints, ok := Process(append([]rune(nil), ri[i:]...))
			if ok {
				total += mul(ints[0], ints[1])
			}
		}
	}

	return total
}

func ProcessDo(ri []rune) bool {
	var i int
	for i = 0; i < len(ri) && i < len(d); i++ {
		if d[i] != ri[i] {
			return false
		}
	}

	if i == len(d) {
		return true
	}

	return false
}

func ProcessDont(ri []rune) bool {
	var i int
	for i = 0; i < len(ri) && i < len(dn); i++ {
		if dn[i] != ri[i] {
			return false
		}
	}

	if i == len(dn) {
		return true
	}

	return false
}

func Process(ri []rune) ([]int, bool) {
	for i := 0; i < len(ri) && i < len(p); i++ {
		// fmt.Println("i, pi", i, pi)
		if p[i] != ri[i] {
			return []int(nil), false
		}
		if ri[i] == '(' {
			ints, err := ProcessInts(append([]rune(nil), ri[i:]...))
			if err != nil {
				continue
			}
			return ints, true
		}
	}

	return []int(nil), false
}

func ProcessInts(ri []rune) ([]int, error) {
	// fmt.Println("got ints to process", string(ri))
	ints := []int{}
	var res string
	for i := range ri {
		if ri[i] == '(' {
			continue
		}
		if len(res) != 0 && ri[i] == ')' {
			if res == "" {
				return ints, errors.New("corrupted")
			}
			v, err := strconv.Atoi(res)
			if err != nil {
				return ints, err
			}
			ints = append(ints, v)
			res = ""

			if len(ints) != 2 {
				return ints, errors.New("corrupted")
			}
			return ints, nil
		}
		if len(res) != 0 && ri[i] == ',' {
			v, err := strconv.Atoi(res)
			if err != nil {
				return ints, err
			}
			ints = append(ints, v)
			// fmt.Println("ints", ints)
			res = ""
			continue
		}
		_, err := strconv.Atoi(string(ri[i]))
		if err != nil {
			return ints, err
		}
		res += string(ri[i])
		// fmt.Println("res", res)
	}

	if len(ints) != 2 {
		return ints, errors.New("corrupted")
	}
	return ints, nil
}

func mul(x, y int) int {
	return x * y
}
