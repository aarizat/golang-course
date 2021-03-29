package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) calculate(input, operator string) int {
	splitInput := strings.Split(input, operator)
	num1 := parse(splitInput[0])
	num2 := parse(splitInput[1])
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}

func parse(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func readStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	return str
}


func main() {
	input := readStdin()
	op := readStdin()
	c := calc{}
	result := c.calculate(input, op)
	fmt.Println(result)
}