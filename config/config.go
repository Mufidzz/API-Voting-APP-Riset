package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/riset_api?parseTime=true&loc=Local")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.User{}, structs.VotingList{}, structs.VotingListItem{})
	return db
}