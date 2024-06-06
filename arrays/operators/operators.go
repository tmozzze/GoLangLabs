package operators

import (
	"arrays/funcs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Sort A+; - отсортировать элементы массива A по неубыванию;
//● Sort A-; - отсортировать элементы массива A по невозрастанию;
//● Shuffle A; - переставить элементы массива в псевдослучайном порядке;
//● Stats a; - вывести в стандартный поток вывода статистическую информацию о массиве
//A: размер массива, максимальный и минимальный элемент (и их индексы), наиболее
//часто встречающийся элемент (если таковых несколько, вывести максимальный из них
//по значению), среднее значение элементов, максимальное из значений отклонений
//элементов от среднего значения;
//● Print a, 3; - вывести в стандартный поток вывода элемент массива A стоящий на позиции
//с номером 3;
//● Print a, 4, 16; - вывести в стандартный поток вывода элементы массива A, начиная с 4 по
//16 включительно оба конца;
//● Print a, all; - вывести в стандартный поток вывод все элементы массива A

func LoadArray(command []string, arrays map[string][]int) {

	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}
	name := command[1]
	fileName := command[2]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 0; i < len(fields); i++ {
			num, err := strconv.Atoi(fields[i])
			if err != nil {
				fmt.Println("Error:", err)
			}

			numbers = append(numbers, num)

		}
	}

	arrays[name] = numbers
	fmt.Printf("Array %s loaded successfully", name)
	fmt.Print("\n")
}

func SaveArray(command []string, arrays map[string][]int) {

	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}
	name := command[1]
	fileName := command[2]

	if !funcs.ArrayExists(name, arrays) {
		return
	}
	numbers := arrays[name]

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Save Error:", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		fmt.Fprintln(writer, num)
	}
	writer.Flush()

	fmt.Println("Array " + name + " saved to " + fileName + " successfully")
}

func RandArray(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]
	count, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Count Error:", err)
	}
	lb, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("Left horizon Error:", err)
	}
	rb, err := strconv.Atoi(command[4])
	if err != nil {
		fmt.Println("Right horizon Error:", err)
	}

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	for i := 0; i < count; i++ {
		num := funcs.RandRange(lb, rb)
		arrays[name] = append(arrays[name], num)
	}
	fmt.Println("Array " + name + " randomized successfully")

}

func ConcatArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}
	nameFirst := command[1]
	nameLast := command[2]

	if !funcs.ArrayExists(nameFirst, arrays) {
		return
	}
	if !funcs.ArrayExists(nameLast, arrays) {
		return
	}

	for _, num := range arrays[nameLast] {
		arrays[nameFirst] = append(arrays[nameFirst], num)
	}

	fmt.Println("Arrays " + nameFirst + " and " + nameLast + " were concatenated successfully")
}

func FreeArray(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}
	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}
	arrays[name] = []int{}

	fmt.Println("Array " + name + " freed successfully")
}

func RemoveElementsFromArray(command []string, arrays map[string][]int) {
	if len(command) != 4 {
		fmt.Println("Invalid command")
		return
	}
	name := command[1]
	index, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Start Error:", err)
	}
	count, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("End Error:", err)
	}

	if index+count > len(arrays[name]) {
		fmt.Println("You want remove too much elements")
		return
	}

	tmp := arrays[name][:index]
	tmp = append(tmp, arrays[name][index+count:]...)
	arrays[name] = tmp

	fmt.Println("Array " + name + " removed successfully")
}

func CopyArray(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command")
		return
	}
	nameFrom := command[1]
	nameTo := command[4]

	if !funcs.ArrayExists(nameFrom, arrays) {
		return
	}
	if !funcs.ArrayExists(nameTo, arrays) {
		return
	}

	start, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Start Error:", err)
	}
	end, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("End Error:", err)
	}
	if len(arrays[nameFrom]) < end {
		fmt.Println("Too much elements")
		return
	}

	var tmp []int
	for i := start; i <= end; i++ {
		tmp = append(tmp, arrays[nameFrom][i])
	}

	arrays[nameTo] = append(arrays[nameTo], tmp...)

	fmt.Println("Array  copied successfully")

}

func Sort(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid Command")
		return
	}

	name := command[1]
	if !funcs.ArrayExists(name, arrays) {
		return
	}

	sign := command[2]
	arrays[name] = funcs.QuickSort(arrays[name], sign)
	fmt.Printf("Arrays \"%s\" was sorted successfully\n", name)
}

func Shuffle(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	arrays[name] = funcs.Mix(arrays[name])
	fmt.Println("Array was shuffled succeessfully")
}

func Stats(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	array := arrays[name]

	length := len(array)
	max, maxIndex := funcs.GetMax(array)
	min, minIndex := funcs.GetMin(array)
	mostCommon := funcs.GetMostCommon(array)
	mean := funcs.GetMean(array)
	deviation := funcs.GetMaxDeviation(array)

	fmt.Printf("Length: %v\n", length)
	fmt.Printf("Maximum: %v (index %v)\n", max, maxIndex)
	fmt.Printf("Minimum: %v (index %v)\n", min, minIndex)
	fmt.Printf("Most common: %v\n", mostCommon)
	fmt.Printf("Mean: %v\n", mean)
	fmt.Printf("Max deviation: %v\n", deviation)
}

func Print(command []string, arrays map[string][]int) {
	length := len(command)
	if length < 3 || length > 4 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	if length == 3 {
		if strings.ToLower(command[2]) == "all" {
			if len(arrays[name]) == 0 {
				fmt.Printf("Array \"%s\" is empty\n", name)
				return
			}

			fmt.Println(arrays[name])
			return
		}

		index, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Println("Invalid command")
			return
		}
		if index >= len(arrays[name]) {
			fmt.Println("Index is out of range")
			return
		}
		fmt.Println(arrays[name][index])
		return
	} else if length == 4 {
		firstIndex, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Println("Bad first index")
			return
		}

		secondIndex, err := strconv.Atoi(command[3])
		if err != nil {
			fmt.Println("Bad second index")
			return
		}

		if firstIndex > secondIndex {
			fmt.Println("Second index is lesser than first")
			return
		}
		length = len(arrays[name]) - 1
		if firstIndex > length || secondIndex > length {
			fmt.Println("First and/or second index out of range")
			return
		}

		fmt.Println(arrays[name][firstIndex : secondIndex+1])
		return
	} else {
		fmt.Println("Invalid command")
		return
	}
}
