package entity

import "tomatoPaper/common/util"

// Admin 用户模型对象
type Admin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        //ID
	PostId     int        `gorm:"column:post_id;comment:'岗位id'" json:"postId"`                                 // 岗位id
	DeptId     int        `gorm:"column:dept_id;comment:'部门id'" json:"deptId"`                                 // 部门id
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`           // 密码
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`                    // 昵称
	Status     int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`                           //  头像
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`                          // 邮箱
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`                          // 电话
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`                           // 备注
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
}

// AddAdminDto 新增参数
type AddAdminDto struct {
	PostId   int    `validate:"required"` // 岗位id
	RoleId   uint   `validate:"required"` // 角色id
	DeptId   int    `validate:"required"` // 部门id
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Nickname string `validate:"required"` // 昵称
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 邮箱
	Note     string // 备注
	Status   int    `validate:"required"` // 状态：1->启用,2->禁用
}
