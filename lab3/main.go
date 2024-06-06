package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Scope struct {
	vars   map[string]string
	parent *Scope
}

// новая область
func NewScope(parent *Scope) *Scope {
	return &Scope{
		vars:   make(map[string]string),
		parent: parent,
	}
}

// переменная в тек. области
func (s *Scope) SetVar(name, value string) {
	s.vars[name] = value
}

// значение переменной
func (s *Scope) GetVar(name string) (string, error) {
	if value, exists := s.vars[name]; exists {
		return value, nil
	} else if s.parent != nil {
		return s.parent.GetVar(name)
	}
	return "", fmt.Errorf("Variable %s not found", name)
}

// выводит переменные в области
func (s *Scope) ShowVar() {
	fmt.Println(s.vars)
}

func (s *Scope) Interpreter(filename string) error {
	globalScope := NewScope(nil)
	currentScope := globalScope
	scopeStack := []*Scope{globalScope}

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, ";")
		for _, command := range commands {
			command = strings.TrimSpace(command)
			if command == "{" {
				newScope := NewScope(currentScope)
				scopeStack = append(scopeStack, newScope)
				currentScope = newScope
			} else if command == "}" {
				if len(scopeStack) > 1 {
					scopeStack = scopeStack[:len(scopeStack)-1]
					currentScope = scopeStack[len(scopeStack)-1]
				}
			} else if command == "ShowVar" {
				currentScope.ShowVar()
			} else if strings.Contains(command, "=") {
				parts := strings.Split(command, "=")
				varName := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				currentScope.SetVar(varName, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Ошибка при чтении файла: %v", err)
	}

	return nil
}

func main() {
	inter := NewScope(nil)
	err := inter.Interpreter("instruction.txt")
	if err != nil {
		fmt.Println(err)
	}
}
