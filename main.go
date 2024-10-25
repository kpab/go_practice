package main

import "fmt"

// 型
// 整数型

func main(){
	var i int = 100
	fmt.Println(i + 50)

	var i2 int64 = 200
	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Print(i + int(i2))
}
