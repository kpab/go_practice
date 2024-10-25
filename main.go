package main

import (
	"fmt"
)

func main(){
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 3
		s2 string = "Hello World"
	)
	fmt.Println(i2, s2)
}
