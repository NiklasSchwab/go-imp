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

// this examples shows if-then-else with scoping rules, the use of or and negation
func ex01() {
	// declaring a new integer x
	l01 := declaration("x", number(1))

	// declaring a new boolean y
	l02 := declaration("y", boolean(true))

	// cond is (x == 0) || (!y == false)
	// x == 0 evaluates to false, but !y evaluates to false and !y == false evaluates to true
	cond := or(equal(variable("x"), number(0)), equal(negation(variable("y")), boolean(false)))

	// assignment in then-block: x = x + 10 --> this works and will leak!
	then01 := assignment("x", plus(variable("x"), number(10)))
	// declaration in then-block: y = 7 --> this works in the scope, but won't leak!
	then02 := declaration("y", number(7))
	// declaration in then-block: z = false --> the new variable won't leave this scope!
	then03 := declaration("z", boolean(false))

	// ignore else in this example
	else01 := declaration("y", number(7))
	else02 := declaration("x", mult(number(7), variable("y")))
	else03 := declaration("z", number(1))
	else04 := assignment("x", plus(variable("x"), variable("z")))

	// generating the blocks of the if-then-else with helper function
	thenBl := block(generateSeq([]Stmt{then01, then02, then03}))
	elseBl := block(generateSeq([]Stmt{else01, else02, else03, else04}))

	l03 := ifthenelse(cond, thenBl, elseBl)
	// x will be 11 (declaration, plus in if-then-else scope)
	l04 := sPrint(variable("x"))
	// y will still be true (no leaking of the if-then-else-scope)
	l05 := sPrint(variable("y"))
	// z will be undefined (new variables from the inner scope are lost)
	l06 := sPrint(variable("z"))

	prog := generateProg([]Stmt{l01, l02, l03, l04, l05, l06})
	prog.run()
}

// this examples shows if-then-else with scoping rules and the use of and (short-circuit evaluation)
func ex02() {
	// this example is based on ex01

	l01 := declaration("x", number(1))
	l02 := declaration("y", boolean(true))

	// cond is (x == 0) && (!y == false)
	// x == 0 evaluates to false, so the condition will always immediately be false
	// --> y == 0 doesn't work, but the and still works --> short circuit evaluation!
	cond := and(equal(variable("x"), number(0)), equal(variable("y"), number(0)))

	// then-block can be ignored in this example
	then01 := assignment("x", plus(variable("x"), number(10)))
	then02 := declaration("y", number(7))
	then03 := declaration("z", boolean(false))

	// declaration in else-block: y = 7 --> this works in the scope and won't leak
	else01 := declaration("y", number(7))
	// declaration of x (doesn't matter if declaration or assignment in this case): multiply 7 with the newly declared y = 7
	// --> the new result of x will leave the scope, while the new y won't leak
	else02 := declaration("x", mult(number(7), variable("y")))
	else03 := declaration("z", number(1))
	// adding the new variable z = 1 to the existing x --> once again, the new x will leave the scope, but z won't
	else04 := assignment("x", plus(variable("x"), variable("z")))

	thenBl := block(generateSeq([]Stmt{then01, then02, then03}))
	elseBl := block(generateSeq([]Stmt{else01, else02, else03, else04}))

	l03 := ifthenelse(cond, thenBl, elseBl)
	// x will evaluate to 50 (x: 0 -> 49 -> 50)
	l04 := sPrint(variable("x"))
	// y will evaluate to true --> new declaration in else scope doesn't leave that scope
	l05 := sPrint(variable("y"))
	// z is undefined, since the declaration won't leak out of its scope
	l06 := sPrint(variable("z"))

	prog := generateProg([]Stmt{l01, l02, l03, l04, l05, l06})
	prog.run()
}

// this example shows while loops with scoping rules and printing
func ex03() {
	// declaring some integer variables
	l01 := declaration("i", number(0))
	l02 := declaration("j", number(5))

	// condition: i < j --> the state of those variables is of the outer scope and is updated after each iteration
	// 					--> type correct assignments will leave the inner scope, but (re-)declarations won't
	cond := lesser(variable("i"), variable("j"))

	// i += 1 in every loop iteration --> this will break the while loop eventually
	do01 := assignment("i", plus(variable("i"), number(1)))
	// y is re-declared in the do scope --> this won't leak!
	do02 := declaration("j", boolean(true))
	// print both variables in each iteration
	do03 := sPrint(variable("i"))
	do04 := sPrint(variable("j"))

	doB := block(generateSeq([]Stmt{do01, do02, do03, do04}))

	l03 := while(cond, doB)

	prog := generateProg([]Stmt{l01, l02, l03})
	prog.run()
}

// this example won't type check and the evaluation fails
func ex04() {
	l01 := declaration("x", number(4))
	// WRONG TYPE! --> evaluation will fail
	l02 := assignment("x", boolean(false))

	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}

// this example shows the correct re-declaration of variables
func ex05() {
	l01 := declaration("x", number(4))
	// this is okay! re-declaring variables works
	l02 := declaration("x", boolean(false))
	// this will print "false", since the re-declaration happened in the same scope
	l03 := sPrint(variable("x"))

	prog := generateProg([]Stmt{l01, l02, l03})
	prog.run()
}

// this example shows the behaviour of undeclared variables
func ex06() {
	// x was never declared!
	l01 := sPrint(variable("x"))
	prog := generateProg([]Stmt{l01})
	prog.run()

	// this example won't type check and "Undefined" will be printed
}

// this example shows type miss-match with ==
func ex07() {
	l01 := declaration("x", number(4))

	// x is of type integer, but a boolean is expected --> evaluation of the while condition will fail
	// also won't type check
	cond := equal(variable("x"), boolean(true))
	do := block(sPrint(variable("x")))

	l02 := while(cond, do)

	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}

// this example shows type miss-match when re-assigning a variable
func ex08() {
	l01 := declaration("x", number(5))
	// x is of type integer, so it can't be assigned to type boolean!
	l02 := assignment("x", boolean(true))
	prog := generateProg([]Stmt{l01, l02})
	prog.run()
}

// this example shows how to use more complex expressions in declarations
func ex09() {
	l01 := declaration("x", number(5))
	// re-declaration works with another type than the original one! --> x := x < 10 --> true
	l02 := declaration("x", lesser(variable("x"), number(10)))
	// "true" will be printed
	l03 := sPrint(variable("x"))

	prog := generateProg([]Stmt{l01, l02, l03})
	prog.run()
}
