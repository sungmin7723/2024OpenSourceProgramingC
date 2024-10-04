package main

import (
	"fmt"
	"reflect"
)

// ASCII 'A' = 65, 'a' = 97, '0' = 48
func main() {
	// i := 13              //var i int = 13
	// var f float64 = 12.9 //f := 12.9
	// fmt.Printf("value i: %d, value f: %f\n", i, f)
	// fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	// fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))

	c1 := 'z' //90
	c2 := '김' //44608

	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2))

}
