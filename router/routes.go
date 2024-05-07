package router

import (
	"monster/controller"
	"monster/service"

	echo "github.com/labstack/echo/v4"
)

type RouterInterface struct {
	MonsterService service.IMonsterService
}

// GetAllRoutes routes
func GetAllRoutes(e *echo.Echo, routerInterface RouterInterface) *echo.Echo {

	// Routes
	e.GET("/monsters", controller.MonsterInfo(routerInterface.MonsterService))
	e.GET("/monsters/:id", controller.MonsterInfo(routerInterface.MonsterService))
	e.GET("/FightAll", controller.FightAllCombination(routerInterface.MonsterService))
	e.POST("/monster", controller.MonsterCreate(routerInterface.MonsterService))
	return e
}
