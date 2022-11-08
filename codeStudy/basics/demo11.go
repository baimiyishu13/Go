package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "haha"
		e
		f = iota
		j
	)

	const (
		n = iota
		m
	)
	fmt.Println(a, b, c, d, e, f, j)
	fmt.Println(n, m)
}
