package gsql

import (
	"lotesaleagent/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string
	CreateAt string
}

func (user *User) BeforeCreate() {
	// UUID version 4
	user.Id = uuid.NewString()
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05+07")
}
func toGormUser(entityUser *model.User) User {
	return User{Username: entityUser.Username, Password: entityUser.Password}
}
func toModelUser(gsqlUser *User) model.User {
	return model.User{Username: gsqlUser.Username, Password: gsqlUser.Password}
}

type GormUserRepository struct {
	db *gorm.DB
}

func (gormUser *GormUserRepository) FindById(userId string) (*model.User, model.WrapError) {
	user := User{}
	err := gormUser.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, model.NewError(err)
	}
	return new(toModelUser(&user)), nil
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (gormUser *GormUserRepository) Create(user *model.User) model.WrapError {
	localUser := toGormUser(user)
	localUser.BeforeCreate()

	result := gormUser.db.Create(&localUser)
	if result.Error != nil {
		return model.NewError(result.Error)
	}
	return nil
}

func (gormUser *GormUserRepository) Find(user *model.User) (*User, model.WrapError) {
	var foundUser User
	result := gormUser.db.Where("username = ? and password = ?", user.Username, user.Password).First(&foundUser)
	if result.Error != nil {
		return nil, model.NewError(result.Error)
	}
	return &foundUser, nil
}
