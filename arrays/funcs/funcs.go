package funcs

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ArrayExists(name string, arrays map[string][]int) bool {
	_, ok := arrays[name]
	if !ok {
		fmt.Printf("Array \"%s\" not found\n", name)
		return ok
	}
	return ok
}

func RandRange(min int, max int) int {
	return min + rand.Intn(max-min)
}

func QuickSort(array []int, sign string) []int {
	if len(array) < 2 {
		return array
	}

	if sign != "+" && sign != "-" {
		fmt.Println("Bad sort argument")
		return array
	}
	var less, more []int

	op := array[0]
	for _, value := range array[1:] {
		if value <= op {
			less = append(less, value)
		} else {
			more = append(more, value)
		}
	}

	if sign == "+" {
		return append(append(QuickSort(less, sign), op), QuickSort(more, sign)...)
	} else {
		return append(append(QuickSort(more, sign), op), QuickSort(less, sign)...)
	}
}

func Mix(arr []int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand.New(seed)
	length := len(arr) - 1
	for i := length; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func GetMin(arr []int) (int, int) {
	min := math.MaxInt
	var minIndex int
	for i, value := range arr {
		if value < min {
			min = value
			minIndex = i
		}
	}

	return min, minIndex
}

func GetMax(arr []int) (int, int) {
	max := math.MinInt
	var maxIndex int
	for i, value := range arr {
		if value > max {
			max = value
			maxIndex = i
		}
	}

	return max, maxIndex
}

func getSet(arr []int) map[int]int {
	set := make(map[int]int)
	for _, value := range arr {
		set[value]++
	}

	return set
}

func GetMostCommon(arr []int) int {
	set := getSet(arr)
	if len(arr) == len(set) {
		max, _ := GetMax(arr)
		return max
	}

	var maxKey, maxValue int
	for key, value := range set {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}

		if value == maxValue {
			maxKey = int(math.Max(float64(maxKey), float64(key)))
		}
	}

	return maxKey
}

func GetMean(arr []int) float64 {
	var sum float64
	for _, value := range arr {
		sum += float64(value)
	}
	return sum / float64(len(arr))
}

func GetMaxDeviation(arr []int) float64 {
	mean := GetMean(arr)
	var maxDev float64
	for _, value := range arr {
		dev := math.Abs(float64(value) - mean)
		if dev > maxDev {
			maxDev = dev
		}
	}

	return maxDev
}
