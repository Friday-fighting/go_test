package main

import (
	"fmt"
)

func deferLearn() {
	f()
	fmt.Println("从 f 正常返回。")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("在 f 中恢复", r)
		}
	}()
	fmt.Println("正在调用 g 。")
	g(0)
	fmt.Println("从 g 正常返回。")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("在 g 中延迟执行", i)
	fmt.Println("在 g 中打印", i)
	g(i + 1)

}
