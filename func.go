package main

import "fmt"

func f() {
	num := []float64{1, 2, 3, 4, 5}
	defer fmt.Println(avg(num))
	x, y := values()
	fmt.Println(x, y)
	fmt.Println(add(5, 5, 0, 5))

	del := func(x, y int) int {
		return x - y
	}

	fmt.Println(del(5, 4))

	fmt.Print(factorial(3))
}

func avg(num []float64) float64 {
	var total float64 = 0
	// total := 0.0
	for _, value := range num {
		total += value
	}
	return total / float64(len(num))
}

func values() (int, int) {
	return 10, 50
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
