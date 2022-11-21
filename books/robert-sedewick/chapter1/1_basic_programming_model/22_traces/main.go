package main

import (
	"fmt"
	"strings"
)

//  1.1.22 Write a version of BinarySearch that uses the recursive rank() given on
//  page 25 and traces the method calls. Each time the recursive method is called,
//  print the argument values lo and hi, indented by the depth of the recursion.
//  Hint: Add an argument to the recursive method that keeps track of the depth.

/**
public static int rank(int key, int[] a) {
	return rank(key, a, 0, a.length - 1);
}

public static int rank(int key, int[] a, int lo, int hi) {
	 // Index of key in a[], if present, is not smaller than lo
     //                                  and not larger than hi.
     if (lo > hi) return -1;
     int mid = lo + (hi - lo) / 2;
     if      (key < a[mid]) return rank(key, a, lo, mid - 1);
     else if (key > a[mid]) return rank(key, a, mid + 1, hi);
     else                   return mid;
}

*/

func Rank(list []int, search int) int {
	return rank(list, search, 0, len(list)-1, 0)
}

func rank(list []int, search int, lo, hi, depth int) int {
	suffix := ""
	if depth != 0 {
		suffix = "⮑ "
	}
	fmt.Print(strings.Repeat("\t", depth) + suffix)
	fmt.Printf("lo=%d hi=%d\n", lo, hi)

	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2
	switch {
	case search > list[mid]:
		return rank(list, search, mid+1, hi, depth+1)
	case search < list[mid]:
		return rank(list, search, lo, mid-1, depth+1)
	default:
		return list[mid]
	}
}
