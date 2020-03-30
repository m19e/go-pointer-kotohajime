package main // import "ptr"

import "fmt"

// Person は人間を表す構造体。
type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインタ型の変数を定義
	// pがポインタ変数
	// *Personポインタ型
	var p *Person

	p = &Person{
		Name: "果穂",
		Age:  12,
	}
	
	fmt.Printf("変数pに格納されているアドレス :%p\n", p)
}
