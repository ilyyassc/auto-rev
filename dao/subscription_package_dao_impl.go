package dao

import (
	"auto-rev/config"
	"auto-rev/model"
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

func (SubscriptionPackageDaoImpl) CreateSubscriptionPackage(data model.SubscriptionPackages) (e error){
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}