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

	*namePoint = "凛世"

	// *namePointは変数なので値を代入できます
	fmt.Printf("*namePointに凛世を代入後の*namePoint :%v\n", *namePoint)

	// 再代入してもnamePointに格納されている*string型の値(アドレス)自体は変わらない
	fmt.Printf("*namePointに凛世を代入後のnamePoint :%v\n", namePoint)

	// stringへのポインタである*string型の値(nameに格納されている値)を上書きしたのでnameの値も変更される
	fmt.Printf("*namePointに凛世を代入後のname :%v\n", name)
}