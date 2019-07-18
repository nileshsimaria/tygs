package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make([]*int, 10)

	for i, v := range a {
		b[i] = &v
	}
	for _, v := range b {
		fmt.Printf("%d ", *v)
	}
	fmt.Println()
}

/*
 *   This program is priting 9 9 9 9 9 9 9 9 9 9
 *   Why ?
 *
 *   Hint :
 *     Think of iteration variable 'v' of line 12
 *     and short variable declaration
 *
 *   Read language specification of for + range @ https://golang.org/ref/spec?source=post_page---------------------------#For_range
 */
