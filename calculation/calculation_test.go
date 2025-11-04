package calculation_test

import (
	"binary_search/calculation"
	"testing"
)

func TestCalculation_OutOfRange(t *testing.T) {
	tests := []struct {
		name      string
		firstNum  int
		SecondNum int
		FindNum   int
		wantErr   bool
	}{
		{name: "big number", firstNum: 1, SecondNum: 10, FindNum: 11, wantErr: true},
		{name: "small number", firstNum: 1, SecondNum: 10, FindNum: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := calculation.Calculation(tt.firstNum, tt.SecondNum, tt.FindNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
