package model

type MonsterInfo struct {
	Id      uint   `json:"id"`
	Name    string `json:"name" binding:"required,gte=1,lte=255"`
	Attack  uint   `json:"attack" binding:"required"`
	Defense uint   `json:"defense" binding:"required"`
	Hp      uint   `json:"hp" binding:"required"`
	Speed   uint   `json:"speed" binding:"required"`
}
