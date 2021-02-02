package dao

import (
	"auto-rev/model"
)

type SubscriptionDao interface {
	GetSubscriptionByUserId(id string) (model.Subscriptions, error)
}
