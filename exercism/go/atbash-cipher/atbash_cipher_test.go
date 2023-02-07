package atbash

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(Atbash("OMG"))
	fmt.Println(Atbash("O M G"))
	fmt.Println(Atbash("mindblowingly"))
}

func TestAtbash(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Atbash(tc.phrase)
			if actual != tc.expected {
				t.Errorf("Atbash('%s'): expected '%s', actual '%s'", tc.phrase, tc.expected, actual)
			}
		})
	}
}

func BenchmarkAtbash(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Atbash(test.phrase)
		}
	}
}
