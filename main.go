package main

import "fmt"

// 型
// 配列型
// あとから要素数を変更できない
func main(){
	// 明示的
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string =  [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	// 暗示的
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数省略
	arr4 := [...]string{"D", "E"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)	

	// 要素数の操作
	fmt.Println(arr2[0])
	fmt.Println(arr2[2])
	arr2[2] = "X"
	fmt.Println(arr2[2])

	// var arr5 [4]int
	// arr5[5]= arr1
	// fmt.Println(arr5)
	fmt.Println(len(arr1))
}
