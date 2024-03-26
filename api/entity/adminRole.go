package entity

type AdminRole struct {
	RoleId  uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`  // 角色id
	AdminId uint `gorm:"column:admin_id;comment:'用户id';NOT NULL" json:"menuId"` // 用户id
}

func (AdminRole) TableName() string {
	return "admin_role"
}
