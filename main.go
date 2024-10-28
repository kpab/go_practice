package main

import (
	"fmt"
)

// for

func main(){
	/*
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("無限")
	}
		*/

	/*
	// while文みたいなの
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
		*/
	
	/*
	// 古典的for
	for i := 0; i < 10; i++{
		if i == 3{
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	*/

	/*
	// 配列をループ
	arr := [3]int{1, 2, 3}
	for i:=0; i<len(arr); i++{
		fmt.Println(arr[i])
	}
	*/

	/*
	// 範囲式for
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	*/

	// スライス
	/*
	sl := []string{"python", "php", "go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	*/

	// map
	m := map[string]int{"apple": 100, "banana": 200}

	for key, value := range m {
		fmt.Println(key, value)
	}

	
}
