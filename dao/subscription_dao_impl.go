package dao

import (
	"auto-rev/config"
	"auto-rev/model"
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
