package model

import(
	"time"
)

type SubscriptionPackages struct {
	UserId string `json:"user_id"`
	PackageId string `json:"package_id"`
	Remain int `json:"remain"`
	Quota int `json:"quota"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	BaseModel
}

func (SubscriptionPackages) TableName() string {
	return "subscription_packages"
}