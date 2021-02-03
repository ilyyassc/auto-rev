package service

import (
	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"
	"auto-rev/external"
)

var subscriptionDao dao.SubscriptionDao = dao.SubscriptionDaoImpl{}
var packageService PackageService = PackageServiceImpl{}
var subscriptionPackageService SubscriptionPackageService = SubscriptionPackageServiceImpl{}
var subscriptionExternal external.SubscriptionExternal = external.SubscriptionExternalImpl{}

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

	if  err != nil {
		return err
	}

	s.SubscriptionPackageId = sp.Id
	s.Remain = p.Quota
	s.Quota = p.Quota
	s.EndDate = sp.EndDate

	err = sImp.CreateSubscription(s)
	if err != nil{
		return err
	}

	go subscriptionExternal.PostAutoRevSubscription(s)

	return nil
}

func (SubscriptionServiceImpl) CreateSubscription(data model.Subscriptions) (e error) {
	defer config.CatchError(&e)
	return subscriptionDao.CreateSubscription(data)
}