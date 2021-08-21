package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Todo struct {
	ID int	`json:"id"`
	Title string	`json:"title"`
	Status bool	`json:"status"`
}

func main() {
	//创建数据库 create database bubble
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@(localhost)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("Open faild,err: ",err)
		return
	}else {
		fmt.Println("连接mysql成功")
	}
	//绑定模型
	db.AutoMigrate(&Todo{})
	//初始化gin框架
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static","static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/**")
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html",nil)
	})

	v1Group := r.Group("v1")
	{
		//添加代办事项
		v1Group.POST("/todo", func(c *gin.Context){
			//添加代办事项
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			err = db.Create(&todo).Error
			//3.返回响应
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
			}else{
				c.JSON(http.StatusOK,todo)
			}

		})

		//查看所有的代办事项
		v1Group.GET("/todo",func(c *gin.Context){
			var todoList []Todo
			err = db.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
			}else{
				c.JSON(http.StatusOK,todoList)
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo:id",func(c *gin.Context){

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id",func(c *gin.Context){
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{
					"error": "id不存在",
				})
				return
			}
			var todo Todo
			err := db.Where("id = ?",id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
				return
			}
			c.BindJSON(&todo)
			if err = db.Save(&todo).Error;err!=nil {
				c.JSON(http.StatusOK, gin.H{
					"error":err.Error(),
				})
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})
		//删除某一个代办事项
		v1Group.DELETE("/todo/:id",func(c *gin.Context){
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{
					"error": "id不存在",
				})
				return
			}
			if err=db.Where("id =?",id).Delete(Todo{}).Error;err != nil {
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,gin.H{
					id:"delete",
				})
			}
		})
	}
	r.Run(":9090")



}