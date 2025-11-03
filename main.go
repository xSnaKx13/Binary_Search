package main

import (
	"binary_search/calculation"
	"binary_search/convert"
	promptdata "binary_search/promptData"
	"fmt"
)

func main() {
	firstNumStr := promptdata.PromptData("Введите число начала диапазона: ")
	SecondNumStr := promptdata.PromptData("Введите число окончания диапазона: ")
	findNumStr := promptdata.PromptData("Введите искомое число в рамках этого диапазона: ")

	firstNum, SecondNum, FindNum, err := convert.Convert(firstNumStr, SecondNumStr, findNumStr)
	if err != nil {
		promptdata.PrintErr(err)
	}
	result, attempts := calculation.Calculation(firstNum, SecondNum, FindNum)
	fmt.Printf("Искомое число найдено: %d\nПопыток затрачено: %d", result, attempts)
}
