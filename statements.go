package main

import "fmt"

// Statement interface
type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

// Different statements

type Seq [2]Stmt
type Decl struct {
	lhs string
	rhs Exp
}
type Assign struct {
	lhs string
	rhs Exp
}
type While struct {
	cond   Exp
	doStmt Stmt
}
type IfThenElse struct {
	cond     Exp
	thenStmt Stmt
	elseStmt Stmt
}
type Print struct {
	printExp Exp
}

// Pretty prints

func (stmt Seq) pretty() string {
	return stmt[0].pretty() + "; " + stmt[1].pretty()
}
func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}
func (asgn Assign) pretty() string {
	return asgn.lhs + " = " + asgn.rhs.pretty()
}
func (while While) pretty() string {
	return "while " + while.cond.pretty() + " {\n" +
		while.doStmt.pretty() +
		"\n}"
}
func (ite IfThenElse) pretty() string {
	return "if " + ite.cond.pretty() + " {\n" +
		ite.thenStmt.pretty() +
		"\n}" + " else " + "{\n" +
		ite.elseStmt.pretty() +
		"\n}"
}
func (p Print) pretty() string {
	return "print " + p.printExp.pretty()
}

// Evals

func (stmt Seq) eval(s ValState) {
	stmt[0].eval(s)
	stmt[1].eval(s)
}
func (decl Decl) eval(s ValState) {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}
func (asgn Assign) eval(s ValState) {
	v := asgn.rhs.eval(s)
	x := (string)(asgn.lhs)
	if _, exists := s[x]; exists {
		s[x] = v
	} else {
		fmt.Printf("assign eval fail")
	}
}
func (while While) eval(s ValState) {
	for {
		v := while.cond.eval(s)
		if v.flag == ValueBool {
			if v.valB == true {
				while.doStmt.eval(s)
			} else {
				break
			}
		} else {
			fmt.Printf("while eval fail")
			break
		}
	}
}
func (ite IfThenElse) eval(s ValState) {
	v := ite.cond.eval(s)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			ite.thenStmt.eval(s)
		case !v.valB:
			ite.elseStmt.eval(s)
		}
	} else {
		fmt.Printf("if-then-else eval fail")
	}
}
func (p Print) eval(s ValState) {
	v := p.printExp.eval(s)
	fmt.Printf(showVal(v))
}

// Type checks

func (stmt Seq) check(t TyState) bool {
	if !stmt[0].check(t) {
		return false
	}
	return stmt[1].check(t)
}
func (decl Decl) check(t TyState) bool {
	ty := decl.rhs.infer(t)
	if ty == TyIllTyped {
		return false
	}
	x := (string)(decl.lhs)
	t[x] = ty
	return true
}
func (a Assign) check(t TyState) bool {
	x := (string)(a.lhs)
	return t[x] == a.rhs.infer(t)
}
func (while While) check(t TyState) bool {
	return while.cond.infer(t) == TyBool
}
func (ite IfThenElse) check(t TyState) bool {
	return ite.cond.infer(t) == TyBool
}
func (p Print) check(t TyState) bool {
	if p.printExp.infer(t) == TyIllTyped {
		return false
	} else {
		return true
	}
}
