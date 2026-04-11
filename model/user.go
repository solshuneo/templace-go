package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"Username" gorm:"primaryKey"`
	Username string `json:"Password" gorm:"unique;not null"`
	Password string `gorm:"<-:create;<-:update;->:false"`
	CreateAt string
}

func (user *User) BeforeCreate() {
	// UUID version 4
	user.Id = uuid.NewString()
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05+07")
}
