package model

type Roles struct {
	BaseModel
	RoleName string `json:"name" gorm:"column:role_name"`
	RoleCode string `json:"code" gorm:"unique;column:code"`
}

func (Roles) TableName() string {
	return "roles"
}
