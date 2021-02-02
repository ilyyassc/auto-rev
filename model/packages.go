package model

type Packages struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Duration int `json:"duration"`
	Quota int `json:"quota"`
	BaseModel
}

func (Packages) TableName() string {
	return "packages"
}