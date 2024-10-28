package main

import (
	"fmt"
)

// switch
// 式スイッチ

func main(){
	/*
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}
	*/

	/*
	switch n:= 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}
	*/

	/*
	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0yoriookiku4yoritiisai")
	case n > 3 && n < 7:
		fmt.Println("nは3より大きく7より小さい")
	}
	*/

	switch n:= 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
		/*
	case n > 3 && n < 6:
		fmt.Println("nは3より大きく6より小さい")
		*/
	default:
		fmt.Println("I don't know")
	}

}
