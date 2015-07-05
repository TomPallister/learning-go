package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func square(x *float64) {
	*x = *x * *x
}

func one(x *int) {
	*x = 1
}

func zero(x *int) {
	*x = 0
}

func fib(n uint) uint {
	if n == 0 {
		return n
	} else if n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func half(x int) (float64, bool) {
	if x%2 == 0 {
		return float64(float64(x) / 2), true
	}
	return float64(float64(x) / 2), false

}

func sum(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, bool) {
	return 5, true
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
