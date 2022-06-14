package main

import "fmt"

// Value State is a mapping from variable names to values
type ValState map[string]Val

// Value State is a mapping from variable names to types
type TyState map[string]Type

func main() {

	fmt.Printf("\n")

	//ex1()
	//ex2()
	//ex3()
	//ex4()

	//ex5()
	//ex6()
	ex7()
}

// TODO: Beispiele für alle speziellen Fälle
// TODO: Leaking scopes?! -->
//								1. Bei While/If-Then-Else muss eine Kopie des State verwendet werden
//								2. Diese Kopie darf sich ändern
//								3. Wenn die Schleife vorbei ist, wird geprüft:
//									--> Ist die Variable geändert? Ist sie auch im alten State? Passt der Typ?
//								4. Es wird entsprechend der alte State angepasst
//								WICHTIG: Die die Condition der Schleife nutzt den neuen, veränderbaren State!!!
// TODO: (Parser)
