package main

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

	code := generateProg([]Stmt{l01, l02, l03})
	code.run()
}

func ex6() {
	l01 := declaration("x", number(0))
	do := sequence(sPrint(variable("x")), assignment("x", plus(variable("x"), number(1))))
	l02 := while(lesser(variable("x"), number(10)), do)

	code := generateProg([]Stmt{l01, l02})
	code.run()
}

func ex7() {
	l01 := declaration("x", boolean(true))
	doThis_01 := declaration("x", number(1))
	doThis_02 := sPrint(number(42))
	doThis := sequence(doThis_01, doThis_02)
	l02 := while(equal(variable("x"), boolean(true)), doThis)
	l03 := sPrint(variable("x"))

	code := generateProg([]Stmt{l01, l02, l03})
	code.run()
}
