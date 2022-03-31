package exampleStruct

import "fmt"

func ExampleStructTest() {
	s11 := NewS11()
	SetS11(s11)
	fmt.Println(s11.GetLen(), s11.GetS1().GetA(), s11.GetS1().GetB(), s11.GetS1().GetC())

}
