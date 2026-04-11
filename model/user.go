package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       string `gorm:"primaryKey"`
	Username string `json:"Username" gorm:"unique;not null"`
	Password string `json:"Password" gorm:"<-:create;<-:update;->:false"`
	CreateAt string
}

func (user *User) BeforeCreate() {
	// UUID version 4
	user.Id = uuid.NewString()
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05+07")
}

func (user *User) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v", user.Id, user.Username, user.Password, user.CreateAt)
}
