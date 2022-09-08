package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id int) (*model.User, error)
	FindByName(name string) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	var user *model.User

	err := r.db.Where("id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByName(name string) ([]*model.User, error) {
	var users []*model.User

	err := r.db.Where("name ILIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user *model.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Save(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
