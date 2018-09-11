package main

import "fmt"
import "strings"
import "math/rand"
import "sort"

func check_expand_policy() {
	fmt.Println("Check slice resize policy")

	aslice := make([]int, 0, 5)
	c := cap(aslice)
	fmt.Println("initial size: ", c)

	for i := 0; i < 25; i++ {
		aslice = append(aslice, i)

		if cap(aslice) != c {
			c = cap(aslice)
			fmt.Println("resize to: ", c)
		}
	}
}

func check_string_slice() {
	fmt.Println("Check string slice")
	haystack := "the spice must flow"
	idx := strings.Index(haystack[5:], " ")
	fmt.Println("index of second blank of ", haystack, ": ", idx+5)
}

func check_copy_slice() {
	fmt.Println("Check copy slice to slice")
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = int(rand.Int31n(1000))
	}
	sort.Ints(numbers)

	worst := make([]int, 5)
	copy(worst, numbers[1:5])
	fmt.Println("2nd to 4th smallest number of")
	fmt.Println(numbers[:10])
	fmt.Println(worst)
}

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Cpy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	check_expand_policy()
	check_string_slice()
	check_copy_slice()
}
