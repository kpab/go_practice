package main

import "fmt"

// クロージャーの実装
// クロージャーとは: 関数と関数の処理に関する関数外の環境をセットとして閉じ込めたもの
func Later() func(string) string {
	var store string
	return func(next string) string{
		s := store
		store = next
		return s
	}
}

func main(){
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("is"))
	fmt.Println(f("golang"))
}
