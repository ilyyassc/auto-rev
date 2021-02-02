package main

import(
	"auto-rev/config"
	"auto-rev/service"
	"auto-rev/dao"
	"auto-rev/controller"
	"gorm.io/gorm"
)

func main(){
	e := config.InitWeb()
	g := initDb()
	service.SetService(g)
	dao.SetDao(g)

	controller.SetInit(e)
	eg := config.SetJwt(e)
	controller.SetSubscription(eg)
	controller.SetUser(eg,e)

	e.Logger.Fatal(e.Start(":1234"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}