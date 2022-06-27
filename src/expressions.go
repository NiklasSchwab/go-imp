package main

import (
	"strconv"
)

// expression interface
type Exp interface {
	pretty() string
	eval(s ValState) Val
	infer(t TyState) Type
}

// the various different expressions
type Num int
type Bool bool
type Plus [2]Exp
type Mult [2]Exp
type Or [2]Exp
type And [2]Exp
type Negation [1]Exp
type Equal [2]Exp
type Lesser [2]Exp
type Group [1]Exp
type Var string

// methods to pretty print expressions
func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}
func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}

}
func (e Plus) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += "+"
	x += e[1].pretty()
	x += ")"
	return x
}
func (e Mult) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += "*"
	x += e[1].pretty()
	x += ")"
	return x
}
func (e Or) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += " || "
	x += e[1].pretty()
	x += ")"
	return x
}
func (e And) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += " && "
	x += e[1].pretty()
	x += ")"
	return x
}
func (e Negation) pretty() string {
	var ret string
	ret = "("
	ret += "!"
	ret += e[0].pretty()
	ret += ")"
	return ret
}
func (e Equal) pretty() string {
	var ret string
	ret = "("
	ret += e[0].pretty()
	ret += "=="
	ret += e[1].pretty()
	ret += ")"
	return ret
}
func (e Lesser) pretty() string {
	var ret string
	ret = "("
	ret += e[0].pretty()
	ret += "<"
	ret += e[1].pretty()
	ret += ")"
	return ret
}
func (e Group) pretty() string {
	var ret string
	ret = "("
	ret += e[0].pretty()
	ret += ")"
	return ret
}
func (x Var) pretty() string {
	return (string)(x)
}

// methods to evaluate expressions
func (x Num) eval(s ValState) Val {
	// a number evaluates to an integer
	return mkInt((int)(x))
}
func (x Bool) eval(s ValState) Val {
	// a bool evaluates to a boolean
	return mkBool((bool)(x))
}
func (e Plus) eval(s ValState) Val {
	// evaluate both sides, and if both evaluate to integers, sum them
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	// return undefined, if not both sides properly evaluate
	return mkUndefined()
}
func (e Mult) eval(s ValState) Val {
	// multiplying is very similar to plus
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}
func (e Or) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == true:
		// if the first condition is a bool and true, the or always "succeeds" and returns true
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		// if both conditions evaluate to booleans, return the or of these booleans
		return mkBool(b1.valB || b2.valB)
	}
	// otherwise, return undefined
	return mkUndefined()
}
func (e And) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		// if the first condition is a boolean and false, the and can immediately evaluate to false
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		// if both sides evaluate to booleans, return the and of these sides
		return mkBool(b1.valB && b2.valB)
	}
	// otherwise return undefined
	return mkUndefined()
}
func (e Negation) eval(s ValState) Val {
	b := e[0].eval(s)
	if b.flag == ValueBool {
		// if the evaluation resulted in a boolean, return it's negation
		return mkBool(!b.valB)
	}
	// otherwise return undefined
	return mkUndefined()
}
func (e Equal) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	if b1.flag == b2.flag {
		switch b1.flag {
		case ValueInt:
			// if both sides evaluate to integers, return the == of these
			return mkBool(b1.valI == b2.valI)
		case ValueBool:
			// if both sides evaluate to booleans, return the == of these
			return mkBool(b1.valB == b2.valB)
		case Undefined:
			// if both sides evaluate to undefined, return undefined
			return mkUndefined()
		}
	}
	// if the type of both sides is different, or otherwise, return undefined
	return mkUndefined()
}
func (e Lesser) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		// if both sides evaluate to integers, return the lesser (a boolean) of these sides
		return mkBool(n1.valI < n2.valI)
	}
	// otherwise return undefined
	return mkUndefined()
}
func (e Group) eval(s ValState) Val {
	return e[0].eval(s)
}
func (x Var) eval(s ValState) Val {
	// evaluating a variable means looking it up in the value state and returning it
	if v, ok := s[string(x)]; ok {
		switch {
		case v.flag == ValueInt:
			// return an integer, if the value is an integer
			return mkInt(v.valI)
		case v.flag == ValueBool:
			// return a boolean, if the value is a bool
			return mkBool(v.valB)
		}
	}
	// otherwise return undefined
	return mkUndefined()
}

// methods to infer/check types of expressions
func (x Num) infer(t TyState) Type {
	return TyInt
}
func (x Bool) infer(t TyState) Type {
	return TyBool
}
func (e Plus) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		// if both sides infer to integer, return int
		return TyInt
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Mult) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		// if both sides infer to integer, return integer
		return TyInt
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Or) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		// if both sides infer to boolean, return bool
		return TyBool
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e And) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		// if both sides infer to boolean, return bool
		return TyBool
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Negation) infer(t TyState) Type {
	t1 := e[0].infer(t)
	if t1 == TyBool {
		// if the expression infers to boolean, return bool
		return TyBool
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Equal) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == t2 {
		// if both sides infer to boolean, return bool
		return TyBool
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Lesser) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		// if both sides infer to integer, return bool
		return TyBool
	}
	// otherwise return IllTyped
	return TyIllTyped
}
func (e Group) infer(t TyState) Type {
	return e[0].infer(t)
}
func (x Var) infer(t TyState) Type {
	// in order to infer the type of a varibale, its type has to be checked in the state
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		// if the variable has an entry in the type state, return the found type
		return ty
	} else {
		// if the variable was not found in the type state, return IllTyped
		return TyIllTyped
	}
}
