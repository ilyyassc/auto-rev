package controller

import(
	"auto-rev/service"
	"auto-rev/config"
	"auto-rev/model"

	"github.com/labstack/echo/v4"
	"errors" 
)

var subscriptionService service.SubscriptionService = service.SubscriptionServiceImpl{}

func SetSubscription(eg *echo.Group) {
	eg.GET("/subscription", getSubscription)
	eg.POST("/subscription", createSubscription)
}

func getSubscription(c echo.Context) (e error) {
	defer config.CatchError(&e)
	test := []string{"username", "id"}
	key := config.GetClaims(c, test...)
	input, _ := key["id"].(string)

	var result, err = subscriptionService.GetSubscriptionByUserId(input)
	if err == nil {
		if result.Count != 0 {
			return res(c, result)
		}
		return resErr(c, errors.New("Data not found"))
	}
	return resErr(c, err)

}

func createSubscription(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := new(model.SubscribeParams)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	err := subscriptionService.Subscribe(data)
	if err == nil {
		return res(c, data)
	}

	return resErr(c, err)
}