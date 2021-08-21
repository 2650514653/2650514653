package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	ID int
	Name string
	Gender string
	Hobby string
}
type Users struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
type User struct {
	ID int64
	Name string		`gorm:"default:'小王子'"`  //指定默认值
	Age int64
}
func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@(localhost)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("Open faild,err: ",err)
		return
	}else {
		fmt.Println("连接mysql成功")
	}

	//把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	u := User{Age:18}
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空

}