package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// removeContentWithinBrackets удаляет все содержимое между фигурными скобками {}
func removeContentWithinBrackets(line string) string {
	bracketLevel := 0
	result := strings.Builder{}
	reader := strings.NewReader(line)

	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if string(char) == "{" {
			bracketLevel++ // Увеличиваем уровень вложенности
			continue
		}

		if string(char) == "}" {
			if bracketLevel > 0 {
				bracketLevel-- // Уменьшаем уровень вложенности
				continue
			}
		}

		if bracketLevel == 0 {
			result.WriteByte(char) // Добавляем символ в результат
		}
	}

	return result.String()
}

// removeComments удаляет все после знака #
func removeComments(line string) string {
	commentIndex := strings.Index(line, "#")
	if commentIndex >= 0 {
		return strings.TrimSpace(line[:commentIndex])
	}
	return strings.TrimSpace(line)
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Удаляем содержимое в фигурных скобках и комментарии после #
		line = removeContentWithinBrackets(line)
		line = removeComments(line)

		if line != "" {
			fmt.Println(line) // Печать результата
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
