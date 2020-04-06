package main	

import "fmt"

var title string = "So"
func main() {
	fmt.Println("Hi, amam")
	fmt.Println("1 + 1 =",1+1)
	fmt.Println("1.0 + 1.0 =",1.0+1.0)
	fmt.Println(len("Hello world"))
	fmt.Println("HelloWorld"[1])
	fmt.Println(`Hi`)
	fmt.Println("Hello"+"Hello")
	fmt.Println(`Hi`+`Hi`)
	fmt.Println(true && true)
	fmt.Println(false || false)
	fmt.Println(!false)

	var text string = "And a miss you again"
	// var text string
	// text = "Hi"
	fmt.Println(text)

	var x string = "hi"
	var y string = "am"
	fmt.Println(x == y)

	t := "He"
	fmt.Println(t)

	num := 80
	fmt.Println(num)

	fmt.Println(title)
	print()

	const post int = 100
	fmt.Println(post)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func print() {
	fmt.Println("Func print : ", title)
}