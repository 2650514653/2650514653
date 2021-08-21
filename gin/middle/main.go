package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context){
	fmt.Println("index in ")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}
//定义一个中间件m1
func m1(c *gin.Context){
	fmt.Println("m1 in")
	c.Next() //调用后续的处理函数
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out")
}

func main() {

	r := gin.Default()
	//全局注册中间件函数m1,m2
	r.Use(m1,m2)

	r.GET("/index", indexHandler) //m1是中间件

	r.Run(":9090")

}
