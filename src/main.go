package main

import "fmt"

// ValState is a mapping from variable names to values
type ValState map[string]Val

// TyState is a mapping from variable names to types
type TyState map[string]Type

func main() {
	fmt.Printf("\n")

	// run the individual examples
	fib()
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
	ex07()
	ex08()
}
