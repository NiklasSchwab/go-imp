package main

import "fmt"

// Statement interface
type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

// Different statements

type Prog [1]Block
type Block [1]Stmt
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
	cond Exp
	do   Block
}
type IfThenElse struct {
	cond   Exp
	thenBl Block
	elseBl Block
}
type Print struct {
	printExp Exp
}

// Pretty prints
func (prg Prog) pretty() string {
	return prg[0].pretty()
}
func (blck Block) pretty() string {
	return "{\n" + blck[0].pretty() + "\n}"
}
func (stmt Seq) pretty() string {
	return stmt[0].pretty() + ";\n" + stmt[1].pretty()
}
func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}
func (asgn Assign) pretty() string {
	return asgn.lhs + " = " + asgn.rhs.pretty()
}
func (while While) pretty() string {
	return "while " + while.cond.pretty() + while.do.pretty()
}
func (ite IfThenElse) pretty() string {
	return "if " + ite.cond.pretty() + ite.thenBl.pretty() + " else " + ite.elseBl.pretty()
}
func (p Print) pretty() string {
	return "print " + p.printExp.pretty()
}

// Evals
func (prg Prog) eval(s ValState) {
	prg[0].eval(s)
}
func (blck Block) eval(s ValState) {
	blck[0].eval(s)
}
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
	oldVal, exists := s[x]
	if exists && (oldVal.flag == v.flag) {
		s[x] = v
	} else {
		fmt.Printf("assign eval fail")
	}
}
func (while While) eval(s1 ValState) {
	s2 := make(map[string]Val)
	for k, v := range s1 {
		s2[k] = v
	}

	for {
		v := while.cond.eval(s2)
		if v.flag == ValueBool {
			if v.valB == true {
				while.do.eval(s2)
				s2 = s1.update(s2)
			} else {
				s3 := s1.update(s2)
				for k, _ := range s1 {
					s1[k] = s3[k]
				}
				break
			}
		} else {
			fmt.Printf("while eval fail")
			break
		}
	}
}
func (ite IfThenElse) eval(s1 ValState) {
	s2 := make(map[string]Val)
	for k, v := range s1 {
		s2[k] = v
	}

	v := ite.cond.eval(s1)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			ite.thenBl.eval(s2)
		case !v.valB:
			ite.elseBl.eval(s2)
		}
	} else {
		fmt.Printf("if-then-else eval fail")
	}
	s3 := s1.update(s2)
	for k, _ := range s1 {
		s1[k] = s3[k]
	}
}
func (p Print) eval(s ValState) {
	v := p.printExp.eval(s)
	fmt.Printf("%s\n", showVal(v))
}

// Type checks
func (prg Prog) check(t TyState) bool {
	return prg[0].check(t)
}
func (blck Block) check(t TyState) bool {
	return blck[0].check(t)
}
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
	if while.cond.infer(t) != TyBool {
		return false
	} else if !while.do.check(t) {
		return false
	} else {
		return true
	}
}
func (ite IfThenElse) check(t TyState) bool {
	if ite.cond.infer(t) != TyBool {
		return false
	} else if !ite.thenBl.check(t) {
		return false
	} else if !ite.elseBl.check(t) {
		return false
	} else {
		return true
	}
}
func (p Print) check(t TyState) bool {
	if p.printExp.infer(t) == TyIllTyped {
		return false
	} else {
		return true
	}
}

// Helper function to update a state
func (s1 ValState) update(s2 ValState) ValState {
	s3 := make(map[string]Val)
	for k, v := range s1 {
		s3[k] = v
		if (s2[k] != s3[k]) && (s2[k].flag == s3[k].flag) {
			s3[k] = s2[k]
		}
	}
	return s3
}
