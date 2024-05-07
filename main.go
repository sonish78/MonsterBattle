package main

import (
	"fmt"
	"monster/db"
	"monster/repo"
	"monster/router"
	"monster/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.Connect()                                       // initialize db with sqllite using migration to build db from the file
	defer db.Close()                                         // close the connection after finished
	monsterRepo := repo.NewMonsterRepo(db)                   // pass in the db to get the instance with the created db above to query
	monsterService := service.NewMonsterService(monsterRepo) // pass in the instance of the repo with the created db to pass in to the routes to the routes to call the controller with the same db connection
	interfaces := router.RouterInterface{
		MonsterService: monsterService,
	}

	e := echo.New()

	e = router.GetAllRoutes(e, interfaces)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "9000"))) // initiate the server in localhost in port 9000
}
