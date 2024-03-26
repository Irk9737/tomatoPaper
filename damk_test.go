package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
	"tomatoPaper/web"
)

// Product ---> 注意结构体名称 Product 而 创建的表的名称为 Products
type Product struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Number         string    `gorm:"unique" json:"number"`                       // 商品编号(唯一)
	Category       string    `gorm:"type:varchar(256);not null" json:"category"` // 商品类别
	Name           string    `gorm:"type:varchar(20);not null" json:"name"`      // 商品名称
	MadeIn         string    `gorm:"type:varchar(128);not null" json:"made_in"`  // 生产地
	ProductionTime time.Time `json:"production_time"`                            //  生产时间
}

type GormResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

var gormDB *gorm.DB
var gormResponse GormResponse

func init() {
	var err error
	sqlStr := "root:password@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	gormDB, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{}) // 配置项预设了连接池 ConnPool
	if err != nil {
		fmt.Println("连接数据库出现了问题:", err)
		return
	}
}

// gorm 依赖: go get gorm.io/gorm
// 数据库驱动: go get gorm.io/driver/mysql
func TestDamk(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/gorm/insert", gormInsertData)
	server.Start(":8080")
}

// gormInsertData 插入操作
func gormInsertData(c *web.Context) {
	// 捕获异常
	defer func() {
		err := recover()
		if err != nil {
			HandleResponse(c, http.StatusBadRequest, "错误", err)
		}
	}()
	var p Product
	err := c.BindJson(&p)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "参数错误", err)
		return
	}
	gormDB.AutoMigrate(&Product{})
	tx := gormDB.Create(&p)
	if tx.RowsAffected > 0 {
		HandleResponse(c, http.StatusOK, "写入成功", "OK")
		return
	}
	fmt.Printf("insert failed, err:%v\n", err)
	HandleResponse(c, http.StatusBadRequest, "写入失败", tx)
	fmt.Println(tx) // 打印结果
}

func HandleResponse(c *web.Context, code int, msg string, data any) {
	gormResponse.Code = code
	gormResponse.Message = msg
	gormResponse.Data = data
	c.RespJSON(http.StatusOK, gormResponse)
}
