package main

import (
	"fmt"

	"github.com/AMFDPMTE/list"
)

func main() {
	fmt.Println("Hello World")
	l := list.New()
	fmt.Println(l.Serialize())
}
