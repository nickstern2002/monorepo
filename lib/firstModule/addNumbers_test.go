package firstModule

import "testing"

// Thanks ChatGPT
func TestStraightAddition(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, tt := range tests {
		result := straightAddition(tt.num1, tt.num2)
		if result != tt.expected {
			t.Errorf("straightAddition(%d, %d) = %d; want %d", tt.num1, tt.num2, result, tt.expected)
		}
	}
}
