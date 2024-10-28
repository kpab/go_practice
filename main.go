package main

import (
	"fmt"
	"os"
)

// defer(デファー): 関数終了時に行いたい処理を指定

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}
// 複数ある場合は下から実行
func RunDefer(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main(){
	TestDefer()

	/*
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	*/

	RunDefer()

	// リソースの解放処理
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}

