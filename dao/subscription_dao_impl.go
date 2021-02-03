package dao

import (
	"auto-rev/config"
	"auto-rev/model"
	"gorm.io/gorm/clause"
)

type SubscriptionDaoImpl struct{}

func (SubscriptionDaoImpl) GetSubscriptionByUserId(id string) (u model.Subscriptions, e error) {
	defer config.CatchError(&e)
	var data model.Subscriptions
	result := g.Where("user_id = ?", id).Find(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

func (SubscriptionDaoImpl) CreateSubscription(data model.Subscriptions) (e error) {
	defer config.CatchError(&e)
	result := g.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}