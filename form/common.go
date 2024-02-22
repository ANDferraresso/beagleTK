package form

//
func stringInSlice(e string, s []string) bool {
   for _, a := range s {
      if a == e {
         return true
      }
   }
   return false
}

/*
The Go 1.21 stable release in August 2023 introduced a new slices package.

package main

import (
	"slices"
)

func main() {
	numbers := []int{0, 42, 10, 8}
	containsTen := slices.Contains(numbers, 10) // true
	containsTwo := slices.Contains(numbers, 2) // false
}
*/