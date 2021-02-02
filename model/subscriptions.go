package model

type Subscriptions struct {
	UserId string `json:"user_id"`
	SubscriptionPackageId string `json:"subscription_package_id"`
	Remain int `json:"remain"`
	Quota int `json:"quota"`
	EndDate *Timestamp `json:"end_date"`
	BaseModel
}

func (Subscriptions) TableName() string {
	return "subscriptions"
}