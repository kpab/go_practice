package main

import "fmt"

// 型
// 文字列型
// 文字列はバイト配列の集まり

func main(){
	var s string = "Hello golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	
	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	rtest
		test
	`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))

}
