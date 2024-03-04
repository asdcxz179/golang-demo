package api

import (
	"github.com/gin-gonic/gin"
)

type RestFulApi interface {
	MakeApi()
}

var Router *gin.Engine

func init() {
	r := gin.Default()
	Router = r
}

func Api(a RestFulApi) {
	a.MakeApi()
}

func GetRouter() *gin.Engine {
	return Router
}
