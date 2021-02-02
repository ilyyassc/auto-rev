package service

import (
	"auto-rev/model"
)

type SubscriptionService interface {
	GetSubscriptionByUserId(user_id string) (model.Subscriptions, error)
}
