package main

import (
	"fmt"
	"reflect"
)

//print the datatype of the variable
func main() {
	a := 100
	fmt.Printf("%T\n", a)
	fmt.Printf(reflect.TypeOf(a).String())

}
