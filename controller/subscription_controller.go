package controller

import(
	"auto-rev/service"
	"auto-rev/config"
	// "auto-rev/model"

	"github.com/labstack/echo/v4"
	"errors"
)

var subscriptionService service.SubscriptionService = service.SubscriptionServiceImpl{}

func SetSubscription(eg *echo.Group) {
	eg.GET("/subscription", getSubscription)
}

func getSubscription(c echo.Context) (e error) {
	defer config.CatchError(&e)
	id := c.Param("userid")
	// id := c.QueryParam("id")

	var result, err = subscriptionService.GetSubscriptionByUserId(id)
	if err == nil {
		if result.Count != 0 {
			return res(c, result)
		}
		return resErr(c, errors.New("Data not found"))
	}
	return resErr(c, err)

}