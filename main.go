package main

import (
	"fmt"
	"mysql"
	"news"
)

func main() {
	news.Insert()
	mysql.Connect()
	fmt.Println("Hello, World!")
}
