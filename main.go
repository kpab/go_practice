package main

import "fmt"

// 関数を引数にとる関数
func CallFunction(f func()){
	f()
}



func main(){
	CallFunction(func(){
		fmt.Println("I'm a function")
	})	
}
