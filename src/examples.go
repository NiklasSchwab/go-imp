package main

// all examples are written with help of the AST helper functions
// the fibonacci example shows how examples are written and executed

// this example shows the fibonacci calculation
func fib() {
	// declare "lines" of the programm
	l01 := declaration("prev", number(-1))
	l02 := declaration("result", number(1))

	// declare the "do" block of a while loop
	do01 := declaration("sum", plus(variable("prev"), variable("result")))
	do02 := assignment("prev", variable("result"))
	do03 := assignment("result", variable("sum"))
	do04 := sPrint(variable("result"))
	doBlock := block(generateSeq([]Stmt{do01, do02, do03, do04}))

	l03 := while(lesser(variable("result"), number(50)), doBlock)

	// generate a program from multiple "lines" of "code"
	prog := generateProg([]Stmt{l01, l02, l03})

	// run the program
	prog.run()
}

func ex01() {
	l01 := declaration("x", number(1))
	l02 := declaration("y", boolean(true))

	cond := or(equal(variable("x"), number(0)), equal(negation(variable("y")), boolean(false)))
	then01 := assignment("x", plus(variable("x"), number(10)))
	then02 := declaration("y", number(7))
	then03 := declaration("z", boolean(false))
	else01 := declaration("y", number(7))
	else02 := declaration("x", mult(number(7), variable("y")))
	else03 := declaration("z", number(1))
	else04 := assignment("x", plus(variable("x"), variable("z")))
	thenBl := block(generateSeq([]Stmt{then01, then02, then03}))
	elseBl := block(generateSeq([]Stmt{else01, else02, else03, else04}))

	l03 := ifthenelse(cond, thenBl, elseBl)
	l04 := sPrint(variable("x"))
	l05 := sPrint(variable("y"))
	l06 := sPrint(variable("z"))

	prog := generateProg([]Stmt{l01, l02, l03, l04, l05, l06})
	prog.run()
}

func ex02() {
	l01 := declaration("x", number(1))
	l02 := declaration("y", boolean(true))

	cond := and(equal(variable("x"), number(0)), equal(negation(variable("y")), boolean(false)))
	then01 := assignment("x", plus(variable("x"), number(10)))
	then02 := declaration("y", number(7))
	then03 := declaration("z", boolean(false))
	else01 := declaration("y", number(7))
	else02 := declaration("x", mult(number(7), variable("y")))
	else03 := declaration("z", number(1))
	else04 := assignment("x", plus(variable("x"), variable("z")))
	thenBl := block(generateSeq([]Stmt{then01, then02, then03}))
	elseBl := block(generateSeq([]Stmt{else01, else02, else03, else04}))

	l03 := ifthenelse(cond, thenBl, elseBl)
	l04 := sPrint(variable("x"))
	l05 := sPrint(variable("y"))
	l06 := sPrint(variable("z"))

	prog := generateProg([]Stmt{l01, l02, l03, l04, l05, l06})
	prog.run()
}

func ex03() {
	l01 := declaration("i", number(0))
	l02 := declaration("j", number(5))

	cond := lesser(variable("i"), variable("j"))
	do01 := assignment("i", plus(variable("i"), number(1)))
	do02 := declaration("j", boolean(true))
	do03 := sPrint(variable("i"))
	do04 := sPrint(variable("j"))
	doB := block(generateSeq([]Stmt{do01, do02, do03, do04}))

	l03 := while(cond, doB)

	prog := generateProg([]Stmt{l01, l02, l03})
	prog.run()
}

func ex04() {
	l01 := declaration("x", number(4))
	l02 := assignment("x", boolean(false))

	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}

func ex05() {
	l01 := declaration("x", number(4))
	l02 := declaration("x", boolean(false))

	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}

func ex06() {
	l01 := assignment("x", number(4))
	prog := generateProg([]Stmt{l01})
	prog.run()
}

func ex07() {
	l01 := sPrint(variable("x"))
	prog := generateProg([]Stmt{l01})
	prog.run()
}

func ex08() {
	l01 := declaration("x", number(4))
	do := block(sPrint(variable("x")))
	cond := equal(variable("x"), boolean(true))
	l02 := while(cond, do)

	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}
