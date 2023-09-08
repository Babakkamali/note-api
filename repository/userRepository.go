package repository

import (
    "github.com/babakkamali/note-api/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (repo *UserRepository) GetUserByPhoneNumber(phone string) (*models.User, error) {
    var user models.User
    err := repo.Db.Where("phone_number = ?", phone).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
    return repo.Db.Create(user).Error
}