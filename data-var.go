package main // 可執行程式必須使用 main 封包
import "fmt" // 載入內建的 fmt 封包，用來做基本輸入輸出

func main() { // 建立 main 函式，程式的進入點
	// 執行程式時，從這裡開始
	// 輸出訊息到終端 fmt.Println(資料,資料, ...)
	// fmt.Println(3) // int
	// fmt.Println(3.14159) // float64
	// fmt.Println("test") // string
	// fmt.Println(true) // bool
	// fmt.Println('a') // 字符 rune

	// 變數的使用
	var x int
	x = 4
	fmt.Println(x)

	var f float64 = 3.1415
	fmt.Println(f)

	var s string = "hello"
	fmt.Println(s)

	var b bool = true
	fmt.Println(b)

	var r rune = 'c'
	fmt.Println(r)
}
