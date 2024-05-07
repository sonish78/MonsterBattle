package controller

import (
	"log"
	"monster/model"
	"monster/service"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

func FightAllCombination(monsterService service.IMonsterService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var wg sync.WaitGroup
		allMonster := monsterService.MonsterInfo("")
		monsterCount := len(allMonster)
		windStatementSlice := []string{}
		//Looping through to get each and every combination to fight between the monster
		for i := 1; i <= monsterCount; i++ {
			for j := 1; j <= monsterCount; j++ {

				if i != j { // don't fight with the same monster
					wg.Add(1)
					go FightBattle(monsterService, i, j, &windStatementSlice, &wg) // Utilizing go routine to run the process concurrently
				}
			}
		}
		wg.Wait()
		return c.JSON(http.StatusOK, windStatementSlice)
	}
}

func FightBattle(monsterService service.IMonsterService, m1 int, m2 int, windStatementSlice *[]string, wg *sync.WaitGroup) string {
	defer wg.Done()
	monster1 := monsterService.MonsterInfo(strconv.FormatUint(uint64(m1), 10))[0] // Get monster info
	monster2 := monsterService.MonsterInfo(strconv.FormatUint(uint64(m2), 10))[0]

	if m1 == 1 {
		time.Sleep(time.Second * 2)
	}
	var damage uint

	var condition bool = true
	var nextturn string

	// Determine which monster will go according to algorithm
	firstToGo := FirstToGo(monster1, monster2)
	if firstToGo == "" {
		log.Fatalln("Error Fetching First to Go for battle")
	}
	for condition {

		if firstToGo == "monster1" || nextturn == "monster1" {
			// Determine how much damage the monster will make
			damage = CalculateDamage(monster1, monster2)
			if damage > monster2.Hp {
				monster2.Hp = 0
			} else {
				monster2.Hp -= damage
			}
			nextturn = "monster2"

		} else if firstToGo == "monster2" || nextturn == "monster2" {
			damage = CalculateDamage(monster2, monster1)
			if damage > monster1.Hp {
				monster1.Hp = 0
			} else {
				monster1.Hp -= damage
			}
			nextturn = "monster1"
		}

		firstToGo = ""

		if monster1.Hp <= 0 || monster2.Hp <= 0 {
			condition = false
		}
	}

	returnStatement := "Battle between '" + monster1.Name + "' and '" + monster2.Name + "'. Winner is "
	if monster1.Hp <= 0 {

		// As the process is running concurrently its directly updating the address of the slice
		*windStatementSlice = append(*windStatementSlice, returnStatement+"**"+monster2.Name+"**")

	} else if monster2.Hp <= 0 {

		*windStatementSlice = append(*windStatementSlice, returnStatement+"**"+monster1.Name+"**")

	}

	return ""
}

// Determine which monster will go according to algorithm
func FirstToGo(monster1 model.MonsterInfo, monster2 model.MonsterInfo) string {
	if monster1.Speed == monster2.Speed {
		if monster1.Attack > monster2.Attack {
			return "monster1"
		} else {
			return "monster2"
		}
	} else if monster1.Speed > monster2.Speed {
		return "monster1"

	} else if monster2.Speed > monster1.Speed {
		return "monster2"
	}
	return ""
}

func CalculateDamage(monster1 model.MonsterInfo, monster2 model.MonsterInfo) uint {
	var damage uint
	if monster1.Attack == monster2.Attack || monster1.Attack < monster2.Defense {
		damage = 1
	} else {
		damage = monster1.Attack - monster2.Defense
	}

	return damage
}
