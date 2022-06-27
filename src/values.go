package main

// Of what kind is the value? --> ValueInt, ValueBool or Undefined
type Kind int

const (
	ValueInt  Kind = 0
	ValueBool Kind = 1
	Undefined Kind = 2
)

// A value object consists of a flag (Kind), a valI and/or a valB
type Val struct {
	flag Kind
	valI int
	valB bool
}

// Return a new int value object
func mkInt(x int) Val {
	return Val{flag: ValueInt, valI: x}
}

// Return a new bool value object
func mkBool(x bool) Val {
	return Val{flag: ValueBool, valB: x}
}

// Return a new undefined type value object
func mkUndefined() Val {
	return Val{flag: Undefined}
}

// Return the value object's value as pretty string
func showVal(v Val) string {
	var s string
	switch {
	case v.flag == ValueInt:
		s = Num(v.valI).pretty()
	case v.flag == ValueBool:
		s = Bool(v.valB).pretty()
	case v.flag == Undefined:
		s = "Undefined"
	}
	return s
}
