package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング

func main(){
	var s string = "10"

	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i)
}
