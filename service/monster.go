package service

import (
	"monster/model"
	"monster/repo"
)

type IMonsterService interface {
	MonsterInfo(string) []model.MonsterInfo
	MonsterCreate(model.MonsterInfo) error
}

type MonsterService struct {
	Repo repo.IMonsterRepo
}

func NewMonsterService(monsterRepo repo.IMonsterRepo) *MonsterService {
	return &MonsterService{
		Repo: monsterRepo,
	}
}

func (service *MonsterService) MonsterInfo(id string) []model.MonsterInfo {
	monsterInfo := service.Repo.MonsterInfo(id)
	return monsterInfo
}

func (service *MonsterService) MonsterCreate(MonsterInfo model.MonsterInfo) error {
	err := service.Repo.MonsterCreate(MonsterInfo)
	return err
}
