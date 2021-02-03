package external

import (
	"auto-rev/config"
	"auto-rev/model"

	"encoding/json"
)

type SubscriptionExternalImpl struct{}

func (SubscriptionExternalImpl) PostAutoRevSubscription(data model.Subscriptions) (e error) {
	defer config.CatchError(&e)

	method := "post"
	addr := "localhost:1234/healthz"
	
	jsonValue, err := json.Marshal(data)
	if err != nil{
		return err
	}

	baseExternal(jsonValue, method, addr)
	
	return nil
}
