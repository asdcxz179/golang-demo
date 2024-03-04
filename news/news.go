package news

import "fmt"

type News struct {
	Title string
	Body  string
}

func Insert() bool {
	fmt.Println("News")
	return true
}
