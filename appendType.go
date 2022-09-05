package main

import "fmt"

func main() {
	s := make([]interface{}, 0, 10)
	s = AppendType(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(len(s), cap(s))
	s = AppendType(s, "a", "b", "c", "d", "e", "f", "g", "k", "l", "m", "n", -1, -1.1, 1.1, "change", 123, 0)
	fmt.Println(len(s), cap(s))
	for k, v := range s {
		fmt.Println(k, v)
	}
}

func AppendType(slice []interface{}, data ...interface{}) []interface{} {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]interface{}, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

