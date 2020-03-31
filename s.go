package main // import "s"

import "fmt"

func Test(a string) {
	a = "園田" + a
	fmt.Println("aの値(関数内):", a)
}

func TestPt(a *string) { // aはstringのポインタ変数になる
	*a = "小糸" // aポインタが示す変数に入るstringを変更する
	fmt.Println("aの値(関数内):", *a)
}

func main() {
	a := "智代子"
	fmt.Println("Test()呼び出し前のaの値:", a)

	// 呼び出し時に引数としてアドレスを渡す
	TestPt(&a)
	fmt.Println("Test()呼び出し後のaの値:", a)
}
