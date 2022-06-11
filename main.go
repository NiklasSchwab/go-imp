package main

import "fmt"

// Value State is a mapping from variable names to values
type ValState map[string]Val

// Value State is a mapping from variable names to types
type TyState map[string]Type

// Examples

func ex1() {
	ast := plus(mult(number(1), number(2)), number(0))
	runExp(ast)
}

func ex2() {
	ast := and(boolean(false), number(0))
	runExp(ast)
}

func ex3() {
	ast := or(boolean(false), number(0))
	runExp(ast)
}

func ex4() {
	ast := negation(negation(equal(lesser(number(3), number(5)), lesser(number(1), number(10)))))
	runExp(ast)
}

func ex5() {
	l01 := declaration("x", number(2))
	l02 := declaration("y", mult(variable("x"), number(3)))
	l03 := sPrint(variable("y"))

	prog := sequence(l01, sequence(l02, l03))

	runProg(prog)
}

func ex6() {
	l01 := declaration("x", number(0))
	do := sequence(sPrint(variable("x")), assignment("x", plus(variable("x"), number(1))))
	l02 := while(lesser(variable("x"), number(10)), do)

	prog := sequence(l01, l02)
	runProg(prog)
}

func main() {

	fmt.Printf("\n")

	//ex1()
	//ex2()
	//ex3()
	//ex4()
	ex5()
	ex6()
}

// TODO: Print verbessern
// TODO: Einfachere Methode, Beispiele zu erstellen
//TODO: Beispiele für alle speziellen Fälle
