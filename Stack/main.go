package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (value int) {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	value = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

func (s *Stack) Peek() (value int) {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	value = s.elements[len(s.elements)-1]
	return value
}

func main() {
	stack := Stack{}

	// Добавляем элементы в стек
	stack.Push(10)
	stack.Push(15)
	stack.Push(30)

	// Получаем элемент с вершины стека
	topElement := stack.Peek()
	fmt.Println("Top element:", topElement) // Выведет: Top element: 30.5

	// Удаляем элементы из стека
	stack.Pop()
	topElement = stack.Peek()
	fmt.Println("Top element after pop:", topElement) // Выведет: Top element after pop: twenty
}
