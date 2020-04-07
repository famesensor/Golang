package main

import "fmt"

func array() {
	var x [5]int
	num := [5]float64{
		50,
		51,
		52,
		53,
		54,
	}
	for i := 0; i < 5; i++ {
		x[i] = i * 100
		fmt.Println(x[i])
	}

	// for #3 (key, value) , if you not used key you must be used _
	var total float64
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println("Avg : ", total/float64(len(x)))
	fmt.Println(num)
}
