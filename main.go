package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go ORM Tutorial")

	a := App{}
	a.Initialize("test.db")
	a.Run(":8010")
}
