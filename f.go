package main // import "f"

import "fmt"

func main() {
	name := "樹里"
	fmt.Printf("name :%v\n", name)

	namePoint := &name

	// namePointは&nameが格納されてるだけなので
	// stringへのポインタである*string型の値が格納されている
	fmt.Printf("namePoint :%v\n", namePoint)

	// namePointが指している変数は"*namePoint"のように"*"をつけて表す
	fmt.Printf("namePoint :%v\n", *namePoint)
}