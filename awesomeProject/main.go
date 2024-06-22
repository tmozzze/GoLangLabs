package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func priority(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return -1
	}
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func infixToRPN(infix string) (string, error) {
	output := []string{}
	operators := []string{}
	tokens := ""

	i := 0
	for i < len(infix) {
		char := rune(infix[i])
		token := string(char)

		if token == " " {
			i++ // Пропускаем пробелы
			continue
		}

		if isOperator(token) {

			for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
			i++
		} else if token == "(" {
			operators = append(operators, token)
			i++
		} else if token == ")" {

			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}

			if len(operators) == 0 {
				return "", fmt.Errorf("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
			i++
		} else if unicode.IsDigit(char) {

			start := i
			for i < len(infix) && (unicode.IsDigit(rune(infix[i])) || infix[i] == '.') {
				i++
			}
			tokens = infix[start:i]
			output = append(output, tokens) // Добавляем число в output
		} else {
			return "", fmt.Errorf("invalid token: %s", token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return "", fmt.Errorf("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	//Вывод результата ОПН
	fmt.Printf("RPN result: %s ", output)
	return strings.Join(output, " "), nil
}

func EvaluateRPN(expression string) (float64, error) {
	stack := []float64{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("insufficient operands for operator: %s", token)
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64

			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			}

			stack = append(stack, result)

		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid RPN expression")
	}

	return stack[0], nil
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expression := scanner.Text()

		// Преобразуем инфиксное выражение в ОПН
		rpnExpression, err := infixToRPN(expression)
		if err != nil {
			fmt.Println("Error converting to RPN:", err)
			return
		}

		// Вычисляем результат в ОПН
		result, err := EvaluateRPN(rpnExpression)
		if err != nil {
			fmt.Println("Error evaluating RPN:", err)
			return
		}

		fmt.Printf("Result: %f\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
