package calculation

import (
	"errors"
)

func Calculation(firstNum, SecondNum, FindNum int) (int, int, error) {
	if FindNum > SecondNum {
		return 0, 0, errors.New("искомое число больше диапазона")
	}
	if FindNum < firstNum {
		return 0, 0, errors.New("искомое число меньше диапазона")
	}

	size := SecondNum - firstNum
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
			return numb[mid], attempts, nil
		} else if numb[mid] < FindNum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0, 0, errors.New("число не найдено")
}
