package secondModul

import (
	"testing"
)

func TestStraightSubtraction(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{3, 2, 1},
		{0, 0, 0},
		{-1, -1, 0},
		{200, 100, 100},
	}

	for _, tt := range tests {
		result := straightSubtraction(tt.num1, tt.num2)
		if result != tt.expected {
			t.Errorf("straightSubtraction(%d, %d) = %d; want %d", tt.num1, tt.num2, result, tt.expected)
		}
	}
}
