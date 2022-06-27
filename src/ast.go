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

func (prg Prog) run() {
	s := make(map[string]Val)
	t := make(map[string]Type)

	fmt.Printf("\n")
	fmt.Printf("CODE FROM AST:\n")
	fmt.Printf("%s\n\n", prg.pretty())
	fmt.Printf("TYPE CHECK: %t\n\n", prg.check(t))
	fmt.Printf("RUNTIME RESULT:\n")
	prg.eval(s)
	fmt.Printf("\n")
	fmt.Printf("\n**************************\n")
	fmt.Printf("\n")

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

func prog(b Block) Prog {
	return [1]Block{b}
}
func block(s Stmt) Block {
	return [1]Stmt{s}
}
func sequence(x Stmt, y Stmt) Stmt {
	return (Seq)([2]Stmt{x, y})
}
func declaration(lhs string, rhs Exp) Stmt {
	return Decl{lhs, rhs}
}
func assignment(lhs string, rhs Exp) Stmt {
	return Assign{lhs, rhs}
}
func while(cond Exp, do Block) Stmt {
	return While{cond, do}
}
func ifthenelse(cond Exp, th Block, el Block) Stmt {
	return IfThenElse{cond, th, el}
}
func sPrint(s Exp) Stmt {
	return Print{s}
}

// Helper to create a program from multiple "lines" of statements
func generateProg(lines []Stmt) Prog {
	return prog(block(generateSeq(lines)))
}

func generateSeq(lines []Stmt) Stmt {
	if len(lines) > 1 {
		return sequence(lines[0], generateSeq(lines[1:]))
	} else if len(lines) == 1 {
		return lines[0]
	} else {
		panic("ERROR WHILE GENERATING SEQUENCES!")
	}
}
