package main

import "fmt"

// 型
// byte(unit8)型
// 文字列はアスキーコードで表現される

func main(){
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}
