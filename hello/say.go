package hello

import "fmt"

func GetHello(name string) string {
	return "hello " + name
}

func JustPrint() {
	fmt.Println("hello")
}