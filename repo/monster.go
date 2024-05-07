package repo

import (
	"database/sql"
	"log"
	"monster/model"
)

type IMonsterRepo interface {
	MonsterInfo(string) []model.MonsterInfo
	MonsterCreate(model.MonsterInfo) error
}

type MonsterRepo struct {
	DB *sql.DB
}

func NewMonsterRepo(db *sql.DB) IMonsterRepo {
	return &MonsterRepo{DB: db}
}

// Get the monster info depending on the condition from the sqllite db
func (repo *MonsterRepo) MonsterInfo(id string) []model.MonsterInfo {
	whereClause := ""
	if id != "" {
		whereClause = "where id = ?"
	}
	rows, err := repo.DB.Query("SELECT * FROM monsters "+whereClause, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	monstersList := make([]model.MonsterInfo, 0)
	for rows.Next() {
		var monsterList model.MonsterInfo
		if err = rows.Scan(&monsterList.Id, &monsterList.Name, &monsterList.Attack, &monsterList.Defense, &monsterList.Hp, &monsterList.Speed); err != nil {
			log.Fatal(err)
		}

		monstersList = append(monstersList, monsterList)
	}
	return monstersList
}

// Create the monster into the sqllite db
func (repo *MonsterRepo) MonsterCreate(MonsterInfo model.MonsterInfo) error {

	insertQuery := `INSERT INTO monsters(name,attack,defense,hp,speed) VALUES(?,?,?,?,?)`

	statement, err := repo.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = statement.Exec(MonsterInfo.Name, MonsterInfo.Attack, MonsterInfo.Defense, MonsterInfo.Hp, MonsterInfo.Speed)
	if err != nil {
		return err
	}
	return nil
}
