package gsql

import (
	"lotesaleagent/model"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func (gormUser *GormUserRepository) FindById(userId string) (*model.User, model.WrapError) {
	user := model.User{}
	err := gormUser.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, model.NewError(err)
	}
	return &user, nil
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (gormUser *GormUserRepository) Create(user *model.User) model.WrapError {
	user.BeforeCreate()

	result := gormUser.db.Create(user)
	if result.Error != nil {
		return model.NewError(result.Error)
	}
	return nil
}

func (gormUser *GormUserRepository) Find(user *model.User) (*model.User, model.WrapError) {
	var foundUser model.User
	result := gormUser.db.Where("username = ? and password = ?", user.Username, user.Password).First(&foundUser)
	if result.Error != nil {
		return nil, model.NewError(result.Error)
	}
	return &foundUser, nil
}
