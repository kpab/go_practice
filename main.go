package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("表示")

	// // fmt 標準
	// fmt.Print("Hello")
	// // 改行
	// fmt.Println("Hello")

	// // Println系:各々の文字列は半角スペースで区切られ、文字列の最後に改行を追加
	// fmt.Println("Hello", "world!!")
	// fmt.Println("Hello", "world!!")

	// // Printf系:フォーマットを指定
	// fmt.Printf("%s\n", "Hello")
	// fmt.Printf("%#v\n", "Hello")

	// Sprint系:出力ではなくフォーマットした結果を文字列で返す
	// s := fmt.Sprint("Hello")
	// s1 := fmt.Sprintf("%v\n", "Hello")
	// s2 := fmt.Sprintln("Hello")

	// fmt.Println(s)
	// fmt.Println(s1)
	// fmt.Println(s2)

	// Fprint系:書き込み先指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")

}

// func main() {
// 	// コマンドラインのオプション処理
// 	// コマンドラインを処理するサンプル
// 	// go run main.go -n 20 -m message -x
	
// 	// オプションの値を格納する変数の定義
// 	var (
// 		max int
// 		msg string
// 		x bool
// 	)
	
// 	// IntVar 整数のオプション
// 	flag.IntVar(&max, "n", 32, "処理数の最大値")
// 	// StringVar 文字列のオプション
// 	flag.StringVar(&msg, "m", "", "処理メッセージ")
// 	// BoolVar bool型のオプション コマンドラインに与えられたらtrue, なければfalse
// 	flag.BoolVar(&x, "x", false, "拡張オプション")

// 	// コマンドラインをパース
// 	flag.Parse()

// 	fmt.Println("処理数の最大値 = ", max)
// 	fmt.Println("処理メッセージ = ", msg)
// 	fmt.Println("拡張オプション = ", x)
// }
/*
// rand
func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())	
	// 	0.813527291469711
	// 0.5598026045235738
	// 0.6695717783859498

	//  現在の時刻をシードに使った擬似乱数の生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// 0~99 の間の擬似乱数
	fmt.Println(rand.Intn(100))
	// int型の擬似乱数
	fmt.Println(rand.Int())
	// int32型の擬似乱数
	fmt.Println(rand.Int31())
	// int64型の擬似乱数
	fmt.Println(rand.Int63())
	// unit32型
	fmt.Println(rand.Uint32())
}

/*
// math
func main() {
	// 小数点以下の切り捨て切り上げ

	// math.Trunc 数値の正負を問わず、小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// math.Floor 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// math.Ceil 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

	// // 絶対値
	// fmt.Println(math.Abs(100))
	// fmt.Println(math.Abs(-100))	

	// // 累乗を求める
	// fmt.Println(math.Pow(0, 2))
	// fmt.Println(math.Pow(2, 2))

	// // 平方根、立方根
	// fmt.Println(math.Sqrt(2))
	// fmt.Println(math.Cbrt(8))

	// // 最大値、最小値
	// fmt.Println(math.Max(1, 2))
	// fmt.Println(math.Min(1, 2))



	// // 数学的な定数
	// // 円周率
	// fmt.Println(math.Pi)

	// // 2の平方根
	// fmt.Println(math.Sqrt2)

	// // 数値型に関する定数
	// // float32で表現可能な最大値
	// fmt.Println(math.MaxFloat32)
	// // float32で表現可能な0ではない最小値
	// fmt.Println(math.SmallestNonzeroFloat32)
	// // float64で表現可能な最大値
	// fmt.Println(math.MaxFloat64)
	// // float64で表現可能な0ではない最小値
	// fmt.Println(math.SmallestNonzeroFloat64)
	// // int64 ver
	// fmt.Println(math.MaxInt64)
	// fmt.Println(math.MinInt64)
}

/*
// time
func main() {
	// 指定時間のスリープ
	// 5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")


	// 時刻の差分を取得
	// t5 := time.Date(2024, 11, 6, 0, 0, 0, 0, time.Local)
	// t6 := time.Now()
	// fmt.Println(t6)

	// // t5 -t6
	// d2 := t5.Sub(t6)
	// fmt.Println(d2)

	// // 時刻を比較する
	// fmt.Println(t6.Before(t5))
	// fmt.Println(t6.After(t5))

	// fmt.Println(t5.Before(t6))
	// fmt.Println(t5.After(t6))

	// 時間の感覚を表現
	// time.Duration型を返す
	// fmt.Println(time.Hour)
	// fmt.Printf("%T\n", time.Hour)
	// fmt.Println(time.Minute)
	// fmt.Println(time.Second)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Microsecond)
	// fmt.Println(time.Nanosecond)

	// // time.Duration型を文字列から生成する
	// d, _ := time.ParseDuration("2h30m")
	// fmt.Println(d)

	// t3 := time.Now()
	// t3 = t3.Add(2*time.Minute + 15*time.Second)
	// fmt.Println(t3)
	// ↑現在時刻の2分15秒後を表すtime.Time型の取得


	// 時間の生成
	// 今の時間
	// t := time.Now()
	// fmt.Println(t)

	// 指定した時間を生成
	// t2 := time.Date(2024, 11, 6, 10, 10, 10, 0, time.Local)
	// fmt.Println(t2)
	// fmt.Println(t2.Year())
	// fmt.Println(t2.YearDay())
	// fmt.Println(t2.Month())
	// fmt.Println(t2.Weekday())
	// fmt.Println(t2.Day())
	// fmt.Println(t2.Hour())
	// fmt.Println(t2.Minute())
	// fmt.Println(t2.Second())
	// fmt.Println(t2.Nanosecond())
	// fmt.Println(t2.Zone())	
}

/*
import (
	"fmt"
	"log"
	"os"
)


// os

func main() {
	// ファイル操作
	// -- OpenFile --
	// O_RDONRY 読み取り専用
	// O_WRONRY 書き込み専用
	// O_RDWR 読み書き可能
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする

	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, _ := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))


	// -- Read --
	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// bs := make([]byte, 128)

	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))
	
	// bs2 := make([]byte, 128)

	// nn, err := f.ReadAt(bs2, 10)
	// fmt.Println(nn)
	// fmt.Println(string(bs2))


	// -- Create --
	// f, _ := os.Create("foo.txt")

	// f.Write([]byte("Hello\n"))

	// f.WriteAt([]byte("Golang"), 6)

	// // ファイルの末尾にオフセットを移動
	// f.Seek(0, io.SeekEnd)

	// f.WriteString("Yaah")


	// -- Open --
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err) 
	// }

	// defer f.Close()

	// Args
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// // os Argsの要素数を表示
	// fmt.Printf("length=%d\n", len(os.Args))
	// // os Args の中身をすべて表示
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// log.Fatal
	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }


	// Exit
	// os.Exit(1)
	// fmt.Println("Start")
	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)
}
	*/