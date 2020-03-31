package main // import "s"

import "fmt"

func Test(a string) {
	fmt.Println("aのアドレス(関数内):", &a)
}

func main() {
	a := "智代子"
	fmt.Println("aのアドレス:", &a)
	Test(a)
}
