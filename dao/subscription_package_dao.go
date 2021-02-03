package dao

import (
	"auto-rev/model"
)

type SubscriptionPackageDao interface {
	GetSubscriptionPackageById(string) (model.SubscriptionPackages, error)
	CreateSubscriptionPackage(model.SubscriptionPackages) (model.SubscriptionPackages, error)
}
