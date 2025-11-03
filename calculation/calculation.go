package calculation

func Calculation(firstNum, SecondNum, FindNum int) (int, int) {
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
			break
		} else if numb[mid] < FindNum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return FindNum, attempts
}
