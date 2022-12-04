package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
	a = append(a[:0], a[1:]...)
	fmt.Println(a)

}
