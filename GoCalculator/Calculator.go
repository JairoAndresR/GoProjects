package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {}

func (calc) operate(input string, operator string) int{ //Receiver function, we can call something like calc.operate
	cleanInput := strings.Split(input, operator) //Split
	operator1 := parse(cleanInput[0]) // asign the respective numbers
	operator2 := parse(cleanInput[1])
	switch operator { //Apply the operator and return the value of the operation
	case "+":
		return operator1+operator2
	case "-":
		return operator1-operator2
	case "*":
		return operator1*operator2
	case "/":
		return operator1/operator2
	default:
		return 0
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin) //Scanner declaration
	scanner.Scan() // Scan call
	return scanner.Text() //Return the text
}

func parse(input string) int{
	operator1, _ := strconv.Atoi(input) //Cast values
	return operator1 // Return the number
}

func main() {
	input := readInput() //Read complete operation
	operator := readInput() // Read the operator
	c := calc{} //Create a new instance for calc struct to be able to use operate function
	fmt.Println(c.operate(input, operator)) // Print the result of operate
}