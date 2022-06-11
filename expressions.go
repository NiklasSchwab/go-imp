package main

import "strconv"

// Expression interface
type Exp interface {
	pretty() string
	eval(s ValState) Val
	infer(t TyState) Type
}

// Different expressions

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

// Pretty prints

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
	x += "||"
	x += e[1].pretty()
	x += ")"
	return x
}
func (e And) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += "&&"
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

// Evaluator

func (x Num) eval(s ValState) Val {
	return mkInt((int)(x))
}
func (x Bool) eval(s ValState) Val {
	return mkBool((bool)(x))
}
func (e Plus) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}
func (e Mult) eval(s ValState) Val {
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
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}
func (e And) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}
func (e Negation) eval(s ValState) Val {
	b := e[0].eval(s)
	if b.flag == ValueBool {
		return mkBool(!b.valB)
	}
	return mkUndefined()
}
func (e Equal) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	if b1.flag == ValueBool && b2.flag == ValueBool {
		return mkBool(b1.valB == b2.valB)
	}
	return mkUndefined()
}
func (e Lesser) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkBool(n1.valI < n2.valI)
	}
	return mkUndefined()
}
func (e Group) eval(s ValState) Val {
	return e[0].eval(s)
}
func (x Var) eval(s ValState) Val {
	if v, ok := s[string(x)]; ok {
		switch {
		case v.flag == ValueInt:
			return mkInt(v.valI)
		case v.flag == ValueBool:
			return mkBool(v.valB)
		}
	}
	return mkUndefined()
}

// Type inferencer/checker

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
		return TyInt
	}
	return TyIllTyped
}
func (e Mult) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyInt
	}
	return TyIllTyped
}
func (e Or) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}
func (e And) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}
func (e Negation) infer(t TyState) Type {
	t1 := e[0].infer(t)
	if t1 == TyBool {
		return TyBool
	}
	return TyIllTyped
}
func (e Equal) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyBool && t2 == TyBool {
		return TyBool
	}
	return TyIllTyped
}
func (e Lesser) infer(t TyState) Type {
	t1 := e[0].infer(t)
	t2 := e[1].infer(t)
	if t1 == TyInt && t2 == TyInt {
		return TyBool
	}
	return TyIllTyped
}
func (e Group) infer(t TyState) Type {
	return e[0].infer(t)
}
func (x Var) infer(t TyState) Type {
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		return ty
	} else {
		return TyIllTyped // variable does not exist yields illtyped
	}

}
