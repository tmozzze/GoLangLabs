package main

import (
	"arrays/funcs"
	"arrays/operators"
	"fmt"
	"os"
	"strings"
)

func main() {
	arrays := make(map[string][]int)

	for {
		fmt.Print("Enter command: ")
		input := funcs.GetUserInput()
		command := strings.Fields(input)
		command[0] = strings.ToLower(command[0])

		switch command[0] {

		case "load":
			operators.LoadArray(command, arrays)
		case "save":
			operators.SaveArray(command, arrays)
		case "rand":
			operators.RandArray(command, arrays)
		case "concat":
			operators.ConcatArray(command, arrays)
		case "free":
			operators.FreeArray(command, arrays)
		case "remove":
			operators.RemoveElementsFromArray(command, arrays)
		case "copy":
			operators.CopyArray(command, arrays)
		case "sort":
			operators.Sort(command, arrays)
		case "shuffle":
			operators.Shuffle(command, arrays)
		case "stats":
			operators.Stats(command, arrays)
		case "print":
			operators.Print(command, arrays)
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid command. Try again!")
		}

	}

}
