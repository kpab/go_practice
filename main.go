package main

import "fmt"

// 定数
// 頭文字を大文字にすると他から呼び出せる
const Pi = 3.14
// 複数定義
const (
	URL = "https://google.com"
	SiteName = "google"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 巨大な数値の定義が可能
// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

// iota
const (
	c0 = iota// 連続する整数の連番を生成する
	c1
	c2
)

func main(){
	fmt.Println(Pi)
	// Pi = 3 // 定数は上書きできない
	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big-1)

	fmt.Println(c0, c1, c2)
}
