package main

import "fmt"

type DemoStruct struct {
	name      string
	value     int
	s_pointer *string
}

func CopyStruct (emptyStruct struct{}) any {
	//x := "Original"
	p := emptyStruct
	q := p
	//q.s_pointer = nil
	y := *p
	q = &y
	return q

}

func main() {
	// x := "Original"
	// p := DemoStruct{value: 20, name: "Struct1", s_pointer: &x}
	// q := p
	// q.s_pointer = nil
	// y := *p.s_pointer
	// q.s_pointer = &y
	// *q.s_pointer = "Altered"

	// fmt.Println(*p.s_pointer)
	// fmt.Println(*q.s_pointer)
	// fmt.Println(&p.s_pointer == &q.s_pointer)
	copyStruct := CopyStruct(DemoStruct{})
	fmt.Println(copyStruct)
}
