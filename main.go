package main

import(
	"github.com/ilyyassc/config"
	"github.com/ilyyassc/service"
	"gorm.io/gorm"
	"fmt"
)

func main(){
	e := config.InitWeb()
	g := initDb()
	service.SetService(g)
	dao.SetDao(g)

	controller.SetInit(e)
	eg := config.SetJwt(e)

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