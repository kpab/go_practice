package main

import "fmt"

// 型 & nil
// interface型
// すべての型と互換性がある
// nil はpythonでいうnone

func main(){
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 2
	// 更新はできるが、計算演算はできない
	// fmt.Println(x+3)
	
}
