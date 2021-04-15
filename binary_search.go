package main

import "fmt"

func BinarySearch(s []int, t int) int {
	l := 0
	r := len(s)
	for l < r {
		m := l + (r-l)/2
		if s[m] == t {
			return m
		}
		if s[m] > t {
			r = m
		} else {
			l = m + 1
		}
	}
	return -1 // not found
}

// lowerBound first index of i, such that s[i] >= x
func lowerBound(s []int, t int) int {
	l := 0
	r := len(s)
	for l < r {
		m := l + (r-l)/2
		if s[m] >= t {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// upperBound first index of i, such that s[i] > x
func upperBound(s []int, t int) int {
	l := 0
	r := len(s)
	for l < r {
		m := l + (r-l)/2
		if s[m] > t {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	s := []int{1, 2, 2, 2, 4, 4, 5}
	fmt.Println(lowerBound(s, 2))
	fmt.Println(lowerBound(s, 3))
	fmt.Println(upperBound(s, 2))
	fmt.Println(upperBound(s, 5))
}
