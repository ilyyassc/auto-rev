package model

type SubscriptionPackages struct {
	UserId string `json:"user_id"`
	Remain int `json:"remain"`
	Quota int `json:"quota"`
	StartDate *Timestamp `json:"start_date"`
	EndDate *Timestamp `json:"end_date"`
	BaseModel
}

func (SubscriptionPackages) TableName() string {
	return "subscription_packages"
}