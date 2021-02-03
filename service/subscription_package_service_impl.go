package service

import (
	"time"

	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"
)

var subscriptionPackageDao dao.SubscriptionPackageDao = dao.SubscriptionPackageDaoImpl{}

type SubscriptionPackageServiceImpl struct{}

func (SubscriptionPackageServiceImpl) GetSubscriptionPackageById(id string) (sp model.SubscriptionPackages, e error) {
	defer config.CatchError(&e)
	sp, err := subscriptionPackageDao.GetSubscriptionPackageById(id)
	return sp, err
}

func (SubscriptionPackageServiceImpl) CreateSubscriptionPackage(p model.Packages, user_id string) (sp model.SubscriptionPackages, e error) {
	defer config.CatchError(&e)
	sp = model.SubscriptionPackages{
		UserId: user_id,
		PackageId: p.Id,
		Remain: p.Quota,
		Quota: p.Quota,
		StartDate: time.Now(),
		EndDate: time.Now().AddDate(0, 0, p.Duration),
	}

	sp, err := subscriptionPackageDao.CreateSubscriptionPackage(sp)
	if err != nil{
		return model.SubscriptionPackages{}, err
	}
	return sp, nil

}