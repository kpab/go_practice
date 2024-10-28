package main

import "fmt"

// init: main関数より先に呼ばれる
// 初期化に使う

func init(){
	fmt.Println("init")
}
func init(){
	fmt.Println("init2")
}

func main(){
	fmt.Println("Main")
}

