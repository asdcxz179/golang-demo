package news

import (
	"api"
	"fmt"
	"mysql"
	"time"

	"github.com/gin-gonic/gin"
)

var db = mysql.GetDB()

type News struct {
	Title   string
	Content string
	Date    string
}

type Api struct {
}

func init() {
	a := Api{}
	api.Api(a)
}

func GetNews() []News {
	var news []News
	db.Find(&news)
	return news
}

func (a Api) MakeApi() {
	api.Router.GET("/news", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}

func Insert() bool {
	news := News{Title: "Hello", Content: "World", Date: time.Now().Format("2006-01-02")}
	result := db.Create(&news)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	fmt.Println("最新消息新增成功")
	return true
}
