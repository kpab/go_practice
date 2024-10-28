package main

import "fmt"

// 算術演算子




func main(){
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE") // 文字列の結合もできるよん

	fmt.Println(5 - 1)

	fmt.Println(5 * 4)

	fmt.Println(60 / 4)

	fmt.Println(9 % 4) // 余り

	// 算術演算子と代入
	n := 0
	n += 4
	fmt.Println(n)
	n++ // n = n + 1
	fmt.Println(n)
	n-- // n = n-1
	fmt.Println(n)

	s:="ABC"
	s+="DE"
	fmt.Println(s)

}
