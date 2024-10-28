package main

import (
	"fmt"
)

// 型変換

func main(){
	/*
	var i int = 1
	fl64 := float64(i)
	fmt. Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)
	i2 := int(fl64)
	fmt.Println(i2)
	fmt.Printf("%T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("fl = %T\n", fl)
	fmt.Println(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)
	// float → int: 小数点以下は切り捨て

	*/
	// string to int
	/*
	var s string = "100"
	fmt.Printf("s = %T\n", s)
	i, _ := strconv.Atoi(s) // 二つ目の値を破棄
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)
	*/
	// int to string
	/*
	var i2 = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)
	*/
	// string to byte
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// byte to string
	h2 := string(b)
	fmt.Println(h2)
}
