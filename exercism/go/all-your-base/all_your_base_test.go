package allyourbase

import (
	"reflect"
	"testing"
)

func Benchmark(b *testing.B) {
	N := 10
	b.Run("normal", func(b *testing.B) {
		lst := make([]int, 0)
		for i := 0; i < b.N; i++ {
			for i := 0; i < N; i++ {
				lst = append(lst, i)
			}
		}
	})

	b.Run("normal, again", func(b *testing.B) {
		lst := make([]int, 0)
		for i := 0; i < b.N; i++ {
			for i := 0; i < N; i++ {
				lst = append(lst, i)
			}
		}
	})

	b.Run("reverse", func(b *testing.B) {
		lst := make([]int, 0)
		for i := 0; i < b.N; i++ {
			for i := 0; i < N; i++ {
				lst = append([]int{i}, lst...)
			}
		}
	})

	b.Run("reverse after", func(b *testing.B) {
		lst := make([]int, 0)
		for i := 0; i < b.N; i++ {
			for i := 0; i < N; i++ {
				lst = append(lst, i)
			}
		}

		for i := 0; i < len(lst)/2; i++ {
			lst[i], lst[len(lst)-i-1] = lst[len(lst)-i-1], lst[i]
		}
	})

	b.Run("normal allocated", func(b *testing.B) {
		lst := make([]int, 0, N)
		for i := 0; i < b.N; i++ {
			for i := 0; i < N; i++ {
				lst = append(lst, i)
			}
		}
	})
}

func TestName(t *testing.T) {
	ConvertToBase(10, []int{42}, 2)
}

func TestConvertToBase(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ConvertToBase(tc.inputBase, tc.inputDigits, tc.outputBase)
			if tc.expectedError != "" {
				if err == nil {
					t.Errorf("ConvertToBase(%d, %#v, %d) expected error: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError)
				} else if tc.expectedError != err.Error() {
					t.Errorf("ConvertToBase(%d, %#v, %d)\nexpected error: %q\ngot: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError, err.Error())
				}
			} else if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("ConvertToBase(%d, %#v, %d) = %#v, want:%#v", tc.inputBase, tc.inputDigits, tc.outputBase, actual, tc.expected)
			}
		})
	}
}
