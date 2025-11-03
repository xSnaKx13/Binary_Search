package convert_test

import (
	"binary_search/convert"
	"testing"
)

func TestConvert_ValidNumbers(t *testing.T) {
	numb1 := "1"
	numb2 := "40"
	findNumb := "22"
	_, _, _, err := convert.Convert(numb1, numb2, findNumb)
	if err != nil {
		t.Errorf("Не ожидалась ошибка, но получили: %v", err)
	}
}

func TestConvert_InvalidNumber(t *testing.T) {
	numb1 := "aaa"
	numb2 := "40"
	findNumb := "22"

	_, _, _, err := convert.Convert(numb1, numb2, findNumb)
	if err == nil {
		t.Errorf("Ожидалась ошибка, но ошибки нет")
	}
}
