package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"type:varchar(60)"`
	LastName  string `gorm:"type:varchar(60)"`
	Username  string `gorm:"type:varchar(60)"`
	Email     string `gorm:"type:varchar(200)"`
	Password  string `gorm:"type:text"`
}

type VotingList struct {
	gorm.Model

	Title     string `gorm:"type:varchar(60)"`
	PostDate  time.Time
	CloseDate time.Time
}

type VotingListItem struct {
	gorm.Model

	Parent      uint
	Title       string `gorm:"type:varchar(60)"`
	Description string `gorm:"type:text"`
	Voters      int
}
