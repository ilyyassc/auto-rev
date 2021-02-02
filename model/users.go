package model

type Users struct {
	Email  string `json:"email" gorm:"unique"`
	Pwd    string `json:"pwd"`
	RoleId string `json:"roleId" gorm:"column:role_id"`
	Token  string `json:"token" gorm:"-"`
	BaseModel
}

func (Users) TableName() string {
	return "users"
}
