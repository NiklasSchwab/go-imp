package main

// value "kind" (the type of a value) are expressed as integers: Int value = 0, Bool value = 1, Undefined = 2
type Kind int

const (
	ValueInt  Kind = 0
	ValueBool Kind = 1
	Undefined Kind = 2
)

// value object consist of a flag (Kind) that contains "type" information,
// an integer value and/or a boolean value
type Val struct {
	flag Kind
	valI int
	valB bool
}

// functions to create new value objects
func mkInt(x int) Val {
	return Val{flag: ValueInt, valI: x}
}
func mkBool(x bool) Val {
	return Val{flag: ValueBool, valB: x}
}
func mkUndefined() Val {
	return Val{flag: Undefined}
}

// return the value object's value as pretty string
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
