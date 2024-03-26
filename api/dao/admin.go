package dao

import (
	"time"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	. "tomatoPaper/pkg/database"
)

// GetAdminByUsername 根据用户名查询用户
func GetAdminByUsername(username string) (sysAdmin entity.Admin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// CreateAdmin 新增用户
func CreateAdmin(dto entity.AddAdminDto) bool {
	sysAdminByUsername := GetAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		return false
	}
	sysAdmin := entity.Admin{
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   util.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}
	err := Db.AutoMigrate(&entity.Admin{}, &entity.AdminRole{})
	if err != nil {
		return false
	}

	tx := Db.Create(&sysAdmin)
	sysAdminExist := GetAdminByUsername(dto.Username)
	var entity entity.AdminRole
	entity.AdminId = sysAdminExist.ID
	entity.RoleId = dto.RoleId
	Db.Create(&entity)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
