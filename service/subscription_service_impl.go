package service

import (
	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"
)

var subscriptionDao dao.SubscriptionDao = dao.SubscriptionDaoImpl{}

type SubscriptionServiceImpl struct{}

func (SubscriptionServiceImpl) GetSubscriptionByUserId(id string) (s model.Subscriptions, e error) {
	defer config.CatchError(&e)
	s, err := subscriptionDao.GetSubscriptionByUserId(id)
	return s, err
}
