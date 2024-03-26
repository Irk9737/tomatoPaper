package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

func CreateAdmin(c *web.Context) {
	var dto entity.AddAdminDto
	//_ = c.BindJSON(&dto)
	_ = c.BindJson(&dto)
	service.AdminService().CreateAdmin(c, dto)
}
