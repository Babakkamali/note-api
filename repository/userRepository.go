package repository

import (
    "github.com/babakkamali/note-api/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (repo *UserRepository) GetUserByPhoneNumber(phone string) (*models.User, error) {
    var user models.User
    err := repo.db.Where("phone_number = ?", phone).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
    return repo.db.Create(user).Error
}