package main

import (
	"fmt"
)

// switch
// 型スイッチ


func anything(a interface{}){
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main(){
	anything("aaa")
	anything(1)

	// 型アサーション
	var x interface{} = 3
	i, isInt := x.(int) // 復元
	fmt.Println(i + 2, isInt)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isStr := x.(string); isStr {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")		
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

}

