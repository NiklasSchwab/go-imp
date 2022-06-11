package main

import "fmt"

// Helper functions to build ASTs by hand

func runExp(e Exp) {
	s := make(map[string]Val)
	t := make(map[string]Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(s)))
	fmt.Printf("\n %s", showType(e.infer(t)))
	fmt.Printf("\n")
}

func runProg(prog Stmt) {
	s := make(map[string]Val)
	t := make(map[string]Type)

	fmt.Printf("\n ************************** \n")
	fmt.Printf("\n %s \n", prog.pretty())
	fmt.Println(prog.check(t))
	prog.eval(s)
	fmt.Printf("\n ************************** \n")
}

// Expressions

func number(x int) Exp {
	return Num(x)
}
func boolean(x bool) Exp {
	return Bool(x)
}
func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})
}
func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}
func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}
func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}
func negation(x Exp) Exp {
	return (Negation)([1]Exp{x})
}
func equal(x, y Exp) Exp {
	return (Equal)([2]Exp{x, y})
}
func lesser(x, y Exp) Exp {
	return (Lesser)([2]Exp{x, y})
}
func group(x Exp) Exp {
	return (Group)([1]Exp{x})
}
func variable(x string) Exp {
	return Var(x)
}

// Statements

func sequence(x Stmt, y Stmt) Stmt {
	return (Seq)([2]Stmt{x, y})
}
func declaration(lhs string, rhs Exp) Stmt {
	return Decl{lhs, rhs}
}
func assignment(lhs string, rhs Exp) Stmt {
	return Assign{lhs, rhs}
}
func while(cond Exp, do Stmt) Stmt {
	return While{cond, do}
}
func ifthenelse(cond Exp, th Stmt, el Stmt) Stmt {
	return IfThenElse{cond, th, el}
}
func sPrint(s Exp) Stmt {
	return Print{s}
}
