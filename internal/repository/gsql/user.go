package gsql

import (
	"lotesaleagent/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string
}

func toGormUser(entityUser *model.User) User {
	return User{Username: entityUser.Username, Password: entityUser.Password}
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (gormUser *GormUserRepository) Create(user *model.User) model.WrapError {
	result := gormUser.db.Create(new(toGormUser(user)))
	if result.Error != nil {
		return model.NewError(&result.Error)
	}
	return nil
}

func (gormUser *GormUserRepository) Find(user *model.User) model.WrapError {
	result := gormUser.db.Where("username = ? and password = ?", user.Username, user.Password).First(user)
	if result.Error != nil {
		return model.NewError(&result.Error)
	}
	return nil
}
