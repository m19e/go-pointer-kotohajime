package main // import "ptr"

import "fmt"

// Person は人間を表す構造体。
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "果穂",
		Age:  12,
	}
	fmt.Printf("最初のp :%+v\n", p)

	p2 := p
	p2.Name = "夏葉"
	p2.Age = 20
	// Not p
	fmt.Printf("p2で夏葉に書き換えを行なったはずのp :%+v\n", p)

	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p
	p3.Name = "夏葉"
	p3.Age = 20

	fmt.Printf("p3で夏葉に書き換えを行なったp :%+v\n", p)
}
