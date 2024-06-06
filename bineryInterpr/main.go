package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func interpretInstructions(filePath string) {
	// Мапа для переменных
	vars := make(map[string]int)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())

		// Обработка
		if strings.HasPrefix(command, "in") {
			parts := strings.Fields(command)
			varName := parts[1]
			fmt.Printf("Введите значение для %s: ", varName)
			var input int
			fmt.Scan(&input)
			vars[varName] = input
		} else if strings.HasPrefix(command, "eq") {
			parts := strings.Split(command, ",")
			lhs := strings.Fields(parts[0])[1]
			if len(parts) == 1 {
				// Обрабатываем случай, когда выполняется eq x, x; для инициализации в 0
				vars[lhs] = 0
			} else {
				rhs := strings.TrimSpace(strings.TrimSuffix(parts[1], ";"))
				if rhs == "0" {
					vars[lhs] = int(^uint(vars[lhs])) & 1 // Инвертируем значение
				} else {
					if _, exists := vars[rhs]; !exists {
						vars[rhs] = 0
					}
					if vars[lhs] == vars[rhs] {
						vars[lhs] = 1
					} else {
						vars[lhs] = 0
					}
				}
			}
		} else if strings.HasPrefix(command, "and") {
			parts := strings.Fields(command)
			lhs := parts[1]
			rhs := strings.TrimSuffix(parts[2], ";")
			if _, exists := vars[lhs]; !exists {
				vars[lhs] = 0
			}
			if _, exists := vars[rhs]; !exists {
				vars[rhs] = 0
			}
			vars[lhs] = vars[lhs] & vars[rhs]
		} else if strings.HasPrefix(command, "out") {
			parts := strings.Fields(command)
			varName := strings.TrimSuffix(parts[1], ";")
			fmt.Printf("Результат: %d\n", vars[varName])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

func main() {
	filePath := "instruction.txt"
	interpretInstructions(filePath)
}
