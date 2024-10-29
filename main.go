package main

import "fmt"

// append make len cap
func main(){
	sl := []int{100, 200}
	fmt.Println(sl)
	// append
	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	// make
	sl2 := make([]int, 5)
	fmt.Println(sl2)
	// len
	fmt.Println(len(sl2))
	// cap キャパシティー（容量）
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
}




// スライス
// 宣言, 操作
// 配列と違い要素数指定しない
/*
func main(){
	var sl []int
	fmt.Println(sl)

	// 明示的
	var sl2 []int = []int{100,200}
	fmt.Println(sl2) 

	// 暗黙的
	sl3 := []string{"A","B"}
	fmt.Println(sl3)

	// make関数
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5)

	fmt.Println(sl5[1])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])

	fmt.Println(sl5[:]) // 全て

	fmt.Println(sl5[len(sl5)-1])

	fmt.Println(sl5[1:len(sl5)-1])
}

*/