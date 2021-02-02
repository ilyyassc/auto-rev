package service

import (
	"auto-rev/model"
)

type SubscriptionPackageService interface {
	GetSubscriptionPackageById(id string) (model.SubscriptionPackages, error)
	CreateSubscriptionPackage(p model.Packages, user_id string) (model.SubscriptionPackages, error)
}
