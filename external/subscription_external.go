package external

import (
	"auto-rev/model"
)

type SubscriptionExternal interface {
	PostAutoRevSubscription(model.Subscriptions) error
}