package model

type GetSubscriptionParams struct {
	RoleName string `json:"name" gorm:"column:role_name"`
	RoleCode string `json:"code" gorm:"unique;column:code"`
}