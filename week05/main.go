package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int = 55
	//var i int
	//i = 55
	i := 55
	// i := 55
	f := 3.33
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(f, math.Floor(3.49))
	fmt.Println(strings.Title("subin park"))
	fmt.Printf("i는 %d\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)
}
