package dao

import (
	"auto-rev/config"
	"auto-rev/model"
	"fmt"
)

type SubscriptionPackageDaoImpl struct{}

func (SubscriptionPackageDaoImpl) GetSubscriptionPackageById(id string) (u model.SubscriptionPackages, e error) {
	defer config.CatchError(&e)
	data := model.SubscriptionPackages{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

func (SubscriptionPackageDaoImpl) CreateSubscriptionPackage(data model.SubscriptionPackages) (resp model.SubscriptionPackages, e error){
	defer config.CatchError(&e)
	fmt.Println("voke")
	result := g.Create(&data)
	fmt.Println("block21")
	fmt.Println(result)
	fmt.Println(data.Id)
	fmt.Println("block21")
	if result.Error == nil {
		return data, nil
	}
	return model.SubscriptionPackages{}, result.Error
}