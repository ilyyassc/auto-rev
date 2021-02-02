package dao

import (
	"auto-rev/model"
)

type SubscriptionPackageDao interface {
	GetSubscriptionPackageById(id string) (model.SubscriptionPackages, error)
	CreateSubscriptionPackage(data model.SubscriptionPackages) error
}
