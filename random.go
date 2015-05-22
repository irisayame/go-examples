package main

import "fmt"
import "math/rand"

func main() {
	p := fmt.Print
	pl := fmt.Println

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pl()

	pl(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()

	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	pl()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	pl()
}
