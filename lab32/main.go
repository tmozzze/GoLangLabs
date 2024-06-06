package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interpreter struct {
	variables map[string]string
	functions map[string]Function
}

type Function struct {
	params []string
	body   string
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		variables: make(map[string]string),
		functions: make(map[string]Function),
	}
}

func (interp *Interpreter) Assign(name, value string) {
	interp.variables[name] = value
}

func (interp *Interpreter) Define(name string, params []string, body string) {
	interp.functions[name] = Function{params, body}
}

func (interp *Interpreter) Execute(name string, params []string) int {
	if function, exists := interp.functions[name]; exists {
		values := function.body
		parameters := make(map[string]string)

		for i, param := range function.params {
			parameters[param] = params[i]
		}

		for param, value := range parameters {
			values = strings.ReplaceAll(values, param, value)
		}

		return interp.Calc(interp.PolandNotation(values))
	}
	return 0
}

func (interp *Interpreter) PolandNotation(expression string) []string {
	var result []string
	var stack []string

	for i := 0; i < len(expression); i++ {
		if isAlphaNumeric(expression[i]) {
			var token string
			for ; i < len(expression) && isAlphaNumeric(expression[i]); i++ {
				token += string(expression[i])
			}
			result = append(result, token)
			i--
		} else if expression[i] == '(' {
			stack = append(stack, "(")
		} else if expression[i] == ')' {
			for stack[len(stack)-1] != "(" {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if isOperator(expression[i]) {
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(string(expression[i])) {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, string(expression[i]))
		}
	}

	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result
}

func isAlphaNumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func (interp *Interpreter) Calc(expression []string) int {
	var stack []int
	for _, token := range expression {
		if val, err := strconv.Atoi(token); err == nil {
			stack = append(stack, val)
		} else {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		}
	}
	return stack[len(stack)-1]
}

func (interp *Interpreter) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("No filename provided")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		commands := strings.Split(line, ";")
		for _, command := range commands {
			command = strings.TrimSpace(command)
			if strings.Contains(command, "=") {
				parts := strings.Split(command, "=")
				varName := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				interp.Assign(varName, value)
			} else if strings.Contains(command, ":") {
				parts := strings.Split(command, ":")
				header := strings.TrimSpace(parts[0])
				body := strings.TrimSpace(parts[1])
				headerParts := strings.Split(header, "(")
				funcName := strings.TrimSpace(headerParts[0])
				params := strings.Split(strings.TrimSpace(strings.TrimSuffix(headerParts[1], ")")), ",")
				interp.Define(funcName, params, body)
			} else if strings.Contains(command, "print") {
				for key, value := range interp.variables {
					if strings.Contains(value, "(") {
						funcName := strings.Split(value, "(")[0]
						params := strings.Split(strings.TrimSuffix(strings.Split(value, "(")[1], ")"), ",")
						for i, param := range params {
							params[i] = strings.TrimSpace(param)
						}
						interp.variables[key] = strconv.Itoa(interp.Execute(funcName, params))
					} else if isExpression(value) {
						interp.variables[key] = strconv.Itoa(interp.Calc(interp.PolandNotation(value)))
					}
				}
				fmt.Println(interp.variables, interp.functions)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Ошибка при чтении файла: %v", err)
	}

	return nil
}

func isExpression(value string) bool {
	return strings.ContainsAny(value, "+-*/")
}

func main() {
	inter := NewInterpreter()
	err := inter.Parse("instructions.txt")
	if err != nil {
		fmt.Println(err)
	}
}
