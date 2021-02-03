package model

import(
	"time"
)

type Subscriptions struct {
	UserId string `json:"user_id" gorm:"unique"`
	SubscriptionPackageId string `json:"subscription_package_id"`
	PackageId string `json:"package_id"`
	Remain int `json:"remain"`
	Quota int `json:"quota"`
	EndDate time.Time `json:"end_date"`
	BaseModel
}

func (Subscriptions) TableName() string {
	return "subscriptions"
}