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
	parent    *Interpreter
}

type Function struct {
	params []string
	body   string
}

func NewInterpreter(parent *Interpreter) *Interpreter {
	return &Interpreter{
		variables: make(map[string]string),
		functions: make(map[string]Function),
		parent:    parent,
	}
}

func (interp *Interpreter) Assign(name, value string) {
	interp.variables[name] = value
}

func (interp *Interpreter) Define(name string, params []string, body string) {
	interp.functions[name] = Function{params, body}
}

func (interp *Interpreter) Execute(name string, params []string) (float64, error) {
	function, exists := interp.functions[name]
	if !exists {
		return 0, fmt.Errorf("function %s not defined", name)
	}

	if len(params) != len(function.params) {
		return 0, fmt.Errorf("function %s expects %d parameters, got %d", name, len(function.params), len(params))
	}

	values := function.body
	parameters := make(map[string]string)
	for i, param := range function.params {
		parameters[param] = params[i]
	}

	for param, value := range parameters {
		values = strings.ReplaceAll(values, param, value)
	}

	// Evaluate the expression with substituted parameters
	result, err := interp.Calc(interp.PolandNotation(values))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (interp *Interpreter) PolandNotation(expression string) []string {
	var result []string
	var stack []string
	var token string

	for i := 0; i < len(expression); i++ {
		if isAlphaNumeric(expression[i]) || expression[i] == '.' {
			token += string(expression[i])
		} else {
			if token != "" {
				result = append(result, token)
				token = ""
			}
			if expression[i] == '(' {
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
	}

	if token != "" {
		result = append(result, token)
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

func (interp *Interpreter) Calc(expression []string) (float64, error) {
	var stack []float64
	for _, token := range expression {
		if val, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, val)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("insufficient values in stack for operation %s", token)
			}
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
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("invalid operator %s", token)
			}
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

func (interp *Interpreter) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("no filename provided")
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
						result, err := interp.Execute(funcName, params)
						if err != nil {
							fmt.Println("Error:", err)
						} else {
							interp.variables[key] = strconv.FormatFloat(result, 'f', -1, 64)
						}
					} else if isExpression(value) {
						result, err := interp.Calc(interp.PolandNotation(value))
						if err != nil {
							fmt.Println("Error:", err)
						} else {
							interp.variables[key] = strconv.FormatFloat(result, 'f', -1, 64)
						}
					}
				}
				fmt.Println(interp.variables, interp.functions)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	return nil
}

func isExpression(value string) bool {
	return strings.ContainsAny(value, "+-*/")
}

func main() {
	inter := NewInterpreter(nil)
	err := inter.Parse("instructions.txt")
	if err != nil {
		fmt.Println(err)
	}
}
