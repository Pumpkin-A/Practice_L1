// Разработать программу, которая в рантайме способна определить тип переменной:
//  int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func defineType(v interface{}) {
	fmt.Printf("type of variable is: %v\n", reflect.TypeOf(v))
}

func main() {
	var (
		varInt        int
		varString     string
		varBool       bool
		varChanInt    chan int
		varChanStruct chan struct{}
	)

	defineType(varInt)
	defineType(varString)
	defineType(varBool)
	defineType(varChanInt)
	defineType(varChanStruct)
}
