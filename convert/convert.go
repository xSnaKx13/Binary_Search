package convert

import (
	promptdata "binary_search/promptData"
	"strconv"
)

func Convert(numb1, numb2, findNum string) (int, int, int, error) {
	number1, err := strconv.Atoi(numb1)
	if err != nil {
		promptdata.PrintErr(err)
	}
	number2, err := strconv.Atoi(numb2)
	if err != nil {
		promptdata.PrintErr(err)
	}
	findNumber, err := strconv.Atoi(findNum)
	if err != nil {
		promptdata.PrintErr(err)
	}
	return number1, number2, findNumber, nil
}
