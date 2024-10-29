package main

import "fmt"

// スライス可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range(s){
		n += v
	}
	return n
}

func main(){
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}


/*
// スライス for
func main(){
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)
	/*
	for _, v := range sl{
		fmt.Println(v)
	}
*/
/*
	// 古典的 for
	for i:=0;i<len(sl);i++ {
		fmt.Println(sl[i])
	}
}
*/

/*
// コピー
func main(){
	/*
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl)

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)
*/
/*
	sl := []int{1,2,3,4}
	sl2 := make([]int,5,10)
	fmt.Println(sl2)
	n := copy(sl2, sl) // nはコピーに成功した数
	fmt.Println(n, sl2)
}
*/

/*
// append make len cap
func main(){
	sl := []int{100, 200}
	// append
	fmt.Println(sl)
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

*/

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