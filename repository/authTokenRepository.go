package repository

import (
    "github.com/babakkamali/note-api/models"
    "gorm.io/gorm"
)

type AuthTokenRepository struct {
    db *gorm.DB
}

func NewAuthTokenRepository(db *gorm.DB) *AuthTokenRepository {
    return &AuthTokenRepository{db}
}

// CreateToken creates a new auth token in the database.
func (repo *AuthTokenRepository) CreateToken(authToken *models.AuthToken) error {
    return repo.db.Create(authToken).Error
}

// GetTokenByUserID retrieves an auth token by the associated user ID.
func (repo *AuthTokenRepository) GetTokenByUserID(userID uint) (*models.AuthToken, error) {
    var authToken models.AuthToken
    err := repo.db.Where("user_id = ?", userID).Last(&authToken).Error
    if err != nil {
        return nil, err
    }
    return &authToken, nil
}

// GetTokenByTokenValue retrieves an auth token by the token value itself.
func (repo *AuthTokenRepository) GetTokenByTokenValue(tokenValue string) (*models.AuthToken, error) {
    var authToken models.AuthToken
    err := repo.db.Where("token = ?", tokenValue).Last(&authToken).Error
    if err != nil {
        return nil, err
    }
    return &authToken, nil
}

// DeleteToken deletes an auth token from the database.
func (repo *AuthTokenRepository) DeleteToken(tokenID uint) error {
    return repo.db.Delete(&models.AuthToken{}, tokenID).Error
}

// DeleteTokenByUserID deletes all auth tokens associated with a user ID.
func (repo *AuthTokenRepository) DeleteTokenByUserID(userID uint) error {
    return repo.db.Where("user_id = ?", userID).Delete(&models.AuthToken{}).Error
}

// UpdateToken updates an existing auth token.
func (repo *AuthTokenRepository) UpdateToken(authToken *models.AuthToken) error {
    return repo.db.Save(authToken).Error
}