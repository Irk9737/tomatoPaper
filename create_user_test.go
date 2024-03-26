package main

import (
	"fmt"
	"net/http"
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateUser(t *testing.T) {
	// 设置测试数据库连接
	err := database.SetupDBLink()
	if err != nil {
		t.Fatalf("failed to setup database link: %v", err)
	}

	server := web.NewHTTPServer()
	server.Post("/api/register", controller.CreateUser)
	server.Start(":8080")
}

func TestHello(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/form", func(ctx *web.Context) {
		err := ctx.Req.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	server.Start(":8080")
}

func TestHelloWorld(t *testing.T) {
	server := web.NewHTTPServer()
	server.Get("/user/123", func(ctx *web.Context) {
		ctx.RespJSON(http.StatusOK, entity.Users{
			ID:       123,
			Username: "邓明",
			Password: "damkge",
			Role:     1,
		})
	})

	server.Start(":8080")
}

func TestCreateAdmin(t *testing.T) {
	// 设置测试数据库连接
	err := database.SetupDBLink()
	if err != nil {
		t.Fatalf("failed to setup database link: %v", err)
	}

	server := web.NewHTTPServer()
	//server.Post("/demo/add", controller.CreateDemo)
	server.Post("/admin/add", controller.CreateAdmin)
	server.Start(":8080")
}
