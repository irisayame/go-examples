package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	fmt.Print("split 7")
	fmt.Println(split(7))

	fmt.Print("swap 1 2")
	fmt.Println(swap("1", "2"))
}
