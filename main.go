package main

import (
	promptdata "binary_search/promptData"
	"fmt"
	"strconv"
)

func main() {
	firstNumStr := promptdata.PromptData("Введите первое число: ")
	SecondtNumStr := promptdata.PromptData("Введите второе число: ")
	findNumStr := promptdata.PromptData("Введите искомое число: ")

	firstNum, err := strconv.Atoi(firstNumStr)
	if err != nil {
		fmt.Println(err)
	}
	SecondtNum, err := strconv.Atoi(SecondtNumStr)
	if err != nil {
		fmt.Println(err)
	}

	FindNum, err := strconv.Atoi(findNumStr)
	if err != nil {
		fmt.Println(err)
	}

	size := SecondtNum - firstNum

	numb := make([]int, (size + 1))
	for i := range numb {
		numb[i] = i + firstNum
	}

	left := 0
	right := len(numb) - 1
	attempts := 0

	for left <= right {
		attempts++
		mid := (left + right) / 2
		if numb[mid] == FindNum {
			fmt.Printf("Искомое число найдено: %d\nПопыток: %d", FindNum, attempts)
			break
		} else if numb[mid] < FindNum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
