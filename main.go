package main

import (
	"fmt"
)

func outer(){
	var s4 string = "outer"
	fmt.Println(s4)
}

func main(){
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	//  暗黙的な定義
	// 変数名 := 値
	i2 := 200
	fmt.Println(i2)

	i2 = 250
	fmt.Println(i2)

	outer()

	var s5 string = "Not Use"
	fmt.Println(s5)
}
