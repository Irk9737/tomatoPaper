package service

import (
	"github.com/go-playground/validator/v10"
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/result"
	"tomatoPaper/web"
)

// IAdminService 定义接口
type IAdminService interface {
	CreateAdmin(c *web.Context, dto entity.AddAdminDto)
}

type AdminServiceImpl struct{}

func (a AdminServiceImpl) CreateAdmin(c *web.Context, dto entity.AddAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingNewAdminParameter), result.ApiCode.GetMessage(result.ApiCode.MissingNewAdminParameter))
		return
	}
	bool := dao.CreateAdmin(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.USERNAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return
	}
	result.Success(c, bool)
	return
}

var adminService = AdminServiceImpl{}

func AdminService() IAdminService {
	return &adminService
}
