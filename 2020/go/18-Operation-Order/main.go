package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"log"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", true)

	total := 0
	total2 := 0
	for _, line := range lines {
		sum1, _ := parseOperation(hlp.StringToSlice(line))
		sum2 := calcExpression(hlp.StringToSlice(line), "+*")

		total += sum1
		total2 += sum2

	}
	fmt.Println("Part1", total)
	fmt.Println("Part2", total2)

}
func calcExpression(tokens []string, opPrecedence string) int {
	/* This implementation does not implement composite functions,functions with variable number of arguments, and unary operators. */
	outQueue := []string{}
	opStack := []string{}
	// while there are tokens to be read:
	for i := 0; i < len(tokens); i++ {
		// read a token.
		t := tokens[i]
		if t == " " {
			continue
		}

		// if the token is a number, then:
		// 	push it to the output queue.
		if hlp.StringIsInt(t) {
			outQueue = append(outQueue, t)
			continue
		}

		// else if the token is a function then:
		// 	push it onto the operator stack
		// else if the token is an operator then:
		if strings.Contains("+-*/", t) {
			// 	while ((there is an operator at the top of the operator stack)
			for (len(opStack) > 0) &&
				// 		  and ((the operator at the top of the operator stack has greater precedence)
				((strings.Index(opPrecedence, opStack[0]) < strings.Index(opPrecedence, t)) ||
					// 			  or (the operator at the top of the operator stack has equal precedence and the token is left associative))
					(strings.Index(opPrecedence, opStack[0]) == strings.Index(opPrecedence, t) && true)) &&
				// 		  and (the operator at the top of the operator stack is not a left parenthesis)):
				(opStack[0] != "(") {
				// 		pop operators from the operator stack onto the output queue.
				outQueue = append(outQueue, opStack[0])
				opStack = opStack[1:]
			}
			// 	push it onto the operator stack.
			opStack = append([]string{t}, opStack...)

			continue
		}
		if t == "(" {
			// else if the token is a left parenthesis (i.e. "("), then:
			// 	push it onto the operator stack.
			opStack = append([]string{t}, opStack...)
			continue
		}

		if t == ")" {
			// else if the token is a right parenthesis (i.e. ")"), then:
			for opStack[0] != "(" {
				// 	while the operator at the top of the operator stack is not a left parenthesis:
				// 		pop the operator from the operator stack onto the output queue.
				outQueue = append(outQueue, opStack[0])
				opStack = opStack[1:]

			}
			// 	/* If the stack runs out without finding a left parenthesis, then there are mismatched parentheses. */

			// 	if there is a left parenthesis at the top of the operator stack, then:
			if opStack[0] == "(" {
				// 		pop the operator from the operator stack and discard it
				opStack = opStack[1:]
			}
			// 	if there is a function token at the top of the operator stack, then:
			// 		pop the function from the operator stack onto the output queue.
			continue
		}
	}

	// /* After while loop, if operator stack not null, pop everything to output queue */
	for len(opStack) > 0 {
		outQueue = append(outQueue, opStack[0])
		opStack = opStack[1:]
	}

	// if there are no more tokens to read then:
	// while there are still operator tokens on the stack:
	// 	/* If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses. */
	// 	pop the operator from the operator stack onto the output queue.
	// exit.
	//fmt.Println(outQueue)

	return calcRPN(outQueue)

}
func calcRPN(tokens []string) int {
	operandStack := []int{}

	for len(tokens) > 0 {
		// t := tokens[len(tokens)-1]
		// tokens = tokens[:len(tokens)-1]
		t := tokens[0]
		tokens = tokens[1:]
		if t == " " {
			continue
		}

		if strings.Contains("+-*/", t) {

			sum := 0
			a := operandStack[1]
			b := operandStack[0]
			operandStack = operandStack[2:]

			switch t {
			case "+":
				sum = a + b
			case "*":
				sum = a * b
			case "/":
				sum = a / b
			case "-":
				sum = a - b
			default:
				log.Panic("Unknown operator")
			}
			operandStack = append([]int{sum}, operandStack...)

		} else {
			operandStack = append([]int{hlp.Atoi(t)}, operandStack...)
		}
	}

	return operandStack[0]
}

func parseOperation(tokens []string) (int, int) {
	sum := 0
	operand := "+"

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		j := 0
		val := 0

		switch t {
		case " ":
			continue
		case "+":
			operand = "+"
			continue
		case "*":
			operand = "*"
			continue
		case "(":
			val, j = parseOperation(tokens[i+1:])
			i += j
		case ")":
			return sum, i + 1

		default:
			val = hlp.Atoi(t)
		}

		if operand == "+" {
			sum += val
		}
		if operand == "*" {
			sum *= val
		}

	}
	return sum, -1
}
