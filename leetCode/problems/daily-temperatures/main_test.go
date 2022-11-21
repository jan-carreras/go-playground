package daily_temperatures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:  []int{30, 40, 50, 60},
			output: []int{1, 1, 1, 0},
		},
		{
			input:  []int{30, 60, 90},
			output: []int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.output, DailyTemperaturesBruteforce(tt.input))
	}
}

func TestDailyTemperaturesAgain(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:  []int{30, 40, 50, 60},
			output: []int{1, 1, 1, 0},
		},
		{
			input:  []int{30, 60, 90},
			output: []int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.output, DailyTemperaturesBruteforceRepeat(tt.input))
	}
}

func TestDailyTemperatures2(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:  []int{30, 40, 50, 60},
			output: []int{1, 1, 1, 0},
		},
		{
			input:  []int{30, 60, 90},
			output: []int{1, 1, 0},
		},
		{
			input:  []int{71, 76, 71, 76, 71, 76, 76, 71, 76, 71, 71, 71, 76, 76, 76, 76, 71, 76, 76, 76, 71, 76},
			output: []int{1, 0, 1, 0, 1, 0, 0, 1, 0, 3, 2, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.output, DailyTemperatures(tt.input))
	}
}
