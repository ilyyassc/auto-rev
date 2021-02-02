package service

import (
	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"
)

var subscriptionDao dao.SubscriptionDao = dao.SubscriptionDaoImpl{}
var packageService PackageService = PackageServiceImpl{}
var subscriptionPackageService SubscriptionPackageService = SubscriptionPackageServiceImpl{}

type SubscriptionServiceImpl struct{}

func (SubscriptionServiceImpl) GetSubscriptionByUserId(id string) (s model.Subscriptions, e error) {
	defer config.CatchError(&e)
	s, err := subscriptionDao.GetSubscriptionByUserId(id)
	return s, err
}

func (sImp SubscriptionServiceImpl) Subscribe(data *model.SubscribeParams) (e error) {
	defer config.CatchError(&e)

	s := model.Subscriptions{
		UserId: data.UserId,
	}

	p, err := packageService.GetPackageById(data.PackageId)
	if  err != nil {
		return err
	}

	s.PackageId = p.Id

	sp, err := subscriptionPackageService.CreateSubscriptionPackage(p, data.UserId)

	s.SubscriptionPackageId = sp.Id
	s.Remain = p.Quota
	s.Quota = p.Quota
	s.EndDate = sp.EndDate

	return sImp.CreateSubscription(s)
}

func (SubscriptionServiceImpl) CreateSubscription(data model.Subscriptions) (e error) {
	defer config.CatchError(&e)
	return subscriptionDao.CreateSubscription(data)
}