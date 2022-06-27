package main

import "fmt"

// statement interface
type Stmt interface {
	pretty() string
	eval(s ValState)
	check(t TyState) bool
}

// the various different statements
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

// methods to pretty print statements
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

// methods to evaluate statements
func (prg Prog) eval(s ValState) {
	// evaluating a program means evaluating it's main block
	prg[0].eval(s)
}
func (blck Block) eval(s ValState) {
	// evaluating a block means evaluating it's statement
	blck[0].eval(s)
}
func (stmt Seq) eval(s ValState) {
	// evaluating a sequence means evaluating each statement, one after one
	stmt[0].eval(s)
	stmt[1].eval(s)
}
func (decl Decl) eval(s ValState) {
	// declaring overwrites already existing variables, no matter what type
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}
func (asgn Assign) eval(s ValState) {
	// assign only works, if the variable already exists and the types match
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
	// create a new temporary state is needed for the nested scope
	s2 := make(map[string]Val)
	for k, v := range s1 {
		s2[k] = v
	}

	for {
		v := while.cond.eval(s2)
		if v.flag == ValueBool {
			if v.valB == true {
				// if the while condition is true, evaluate the do block (with the temp state)
				while.do.eval(s2)
				// after evaluating the do block, update state --> this state will "leak"!
				s2 = s1.update(s2)
			} else {
				// if the while condition is false, "break" the while loop
				// now, update the original state, based on the temp state
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
	// create a new temporary state is needed for the nested scope
	s2 := make(map[string]Val)
	for k, v := range s1 {
		s2[k] = v
	}

	// evaluate the condition and then evaluate the block according to the result, or show an error if it failed
	v := ite.cond.eval(s1)
	if v.flag == ValueBool {
		switch {
		case v.valB:
			// evaluate the then block with the temp state
			ite.thenBl.eval(s2)
		case !v.valB:
			// evaluate the else block with the temp state
			ite.elseBl.eval(s2)
		}
	} else {
		fmt.Printf("if-then-else eval fail")
	}

	// after evaluatin the if-then-else, update the original state based on the temp state
	s3 := s1.update(s2)
	for k, _ := range s1 {
		s1[k] = s3[k]
	}
}
func (p Print) eval(s ValState) {
	// evaluating a print means to just print the evaluation result...
	v := p.printExp.eval(s)
	fmt.Printf("%s\n", showVal(v))
}

// methods to type-check statements
func (prg Prog) check(t TyState) bool {
	// type checking a block means checking its "main" block
	return prg[0].check(t)
}
func (blck Block) check(t TyState) bool {
	// type checking a block means checking its inner statement
	return blck[0].check(t)
}
func (stmt Seq) check(t TyState) bool {
	// both statements of a sequence have to successfully type check
	if !stmt[0].check(t) {
		return false
	}
	return stmt[1].check(t)
}
func (decl Decl) check(t TyState) bool {
	// the right-hand-side has to be a correctly typed expression
	ty := decl.rhs.infer(t)
	if ty == TyIllTyped {
		return false
	}
	// remember the variable's type in the state
	x := (string)(decl.lhs)
	t[x] = ty

	return true
}
func (a Assign) check(t TyState) bool {
	// the variable's type in the state has to match the assignment's right-hand-side's type
	x := (string)(a.lhs)
	return t[x] == a.rhs.infer(t)
}
func (while While) check(t TyState) bool {
	// both, condition and do block of the loop, have to successfully type check
	if while.cond.infer(t) != TyBool {
		// condition, then- and else-block all have to successfully type check
		return false
	} else if !while.do.check(t) {
		return false
	} else {
		return true
	}
}
func (ite IfThenElse) check(t TyState) bool {
	// condition, then- and else-block all have to successfully type check
	if ite.cond.infer(t) != TyBool {
		// the condition's type always has to be bool
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
	// the expression to print has to be correctly typed
	if p.printExp.infer(t) == TyIllTyped {
		return false
	} else {
		return true
	}
}

// helper function to update the value state environment
// this is necessary to support nested scopes and prevent unwanted leaking
// returns a value state, which is the updated version of a ValState s1, updated with values from a ValState s2

func (s1 ValState) update(s2 ValState) ValState {
	s3 := make(map[string]Val)

	// only consider values already existing in s1, but each value in s1 is included in the new state
	for k, v := range s1 {
		s3[k] = v

		// update a value if it has changed and the type remains the same
		if (s2[k] != s3[k]) && (s2[k].flag == s3[k].flag) {
			s3[k] = s2[k]
		}
	}
	return s3

	// (we don't have to update the type state, since type changes are unwanted when updating an outer scope)
}
