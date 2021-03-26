package main

import "fmt"

func printMessage(){
	fmt.Println("Hello World!")
}

func threeArgs (a, b, int, s string){
	fmt.Println(a, b, s)
}

func doubleValue(x int) int {
	return x * 2
}

func doubleReturn (x int) (a, b int) {
	return x, x * 2
}

func main() {
	printMessage()
	value := doubleValue(10)
	fmt.Println("Value = ", value)
	value1, value2 := doubleReturn(10)
	fmt.Printf("value1 = %d and value2 = %d\n", value1, value2)
}
