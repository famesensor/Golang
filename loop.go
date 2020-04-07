package main

import "fmt"

func loop() {
	// i := 1
	// for loop
	// #1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// #2
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		if (j % 2) == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}
