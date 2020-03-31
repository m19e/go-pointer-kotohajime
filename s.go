package main // import "s"

import "fmt"

func Test(a string) {
	a = "園田" + a
	fmt.Println("aの値(関数内):", a)
}

func main() {
	a := "智代子"
	fmt.Println("Test()呼び出し前のaの値:", a)
	Test(a)
	fmt.Println("Test()呼び出し後のaの値:", a)
}
