// Разработать программу, которая в рантайме способна определить тип переменной:
//  int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func printType(v interface{}) {
	fmt.Printf("type of variable is: %T\n", v)
}

func defineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("variable type: int")
	case string:
		fmt.Println("variable type: string")
	case bool:
		fmt.Println("variable type: bool")
	case chan int:
		fmt.Println("variable type: channel (chan int)")
	default:
		fmt.Println("variable type: unknown")
	}
}

func main() {
	var (
		varInt        int
		varString     string
		varBool       bool
		varChanInt    chan int
		varChanStruct chan struct{}
	)
	printType(varInt)
	printType(varString)
	printType(varBool)
	printType(varChanInt)
	printType(varChanStruct)

	defineType(varBool)
	defineType(varChanInt)
}
