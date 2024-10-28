package main

import "fmt"

// panic & recover
// 例外処理(pgmを強制終了させる)

func main(){
	defer func(){
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	
	panic("runtime error")
	fmt.Println("START")
}

