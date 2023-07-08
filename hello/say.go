package hello

import "fmt"

func GetHello(name string) string {
	return "hello double " + name
}

func JustPrint() {
	fmt.Println("hello")
}