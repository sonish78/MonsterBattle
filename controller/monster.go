package controller

import (
	"monster/model"
	"monster/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get monster information
func MonsterInfo(monsterService service.IMonsterService) echo.HandlerFunc {
	return func(c echo.Context) error {
		monsterId := c.Param("id")
		monsterList := monsterService.MonsterInfo(monsterId)

		return c.JSON(http.StatusOK, monsterList)
	}
}

// Create monster
func MonsterCreate(monsterService service.IMonsterService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var monsterInput struct {
			Name    string `json:"name"`
			Attack  uint   `json:"attack"`
			Defense uint   `json:"defense"`
			Hp      uint   `json:"hp"`
			Speed   uint   `json:"speed"`
		}
		if err := c.Bind(&monsterInput); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Bind input")
		}

		monsterInfo := model.MonsterInfo{
			Name:    monsterInput.Name,
			Attack:  monsterInput.Attack,
			Defense: monsterInput.Defense,
			Hp:      monsterInput.Hp,
			Speed:   monsterInput.Speed,
		}

		//Validation check -  NOTE :it can be done using godash or gin in simple way but due to time constraint i used the simple technique here
		if monsterInfo.Name == "" || monsterInfo.Attack == 0 || monsterInfo.Defense == 0 || monsterInfo.Hp == 0 || monsterInfo.Speed == 0 {
			return c.String(http.StatusForbidden, "All input required!")
		}

		err := monsterService.MonsterCreate(monsterInfo)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error Creating Monster")
		}

		return c.String(http.StatusOK, "Successfully Created")
	}
}
