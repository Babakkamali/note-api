package services

import (
	"errors"
	"time"

	"github.com/babakkamali/note-api/models"
	"github.com/babakkamali/note-api/repository"
	"github.com/babakkamali/note-api/utils"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) GenerateAndSendSMSToken(phoneNumber string) (string, error) {
    smsToken := utils.GenerateVerificationCode(6)

    user, err := s.userRepo.GetUserByPhoneNumber(phoneNumber)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            newUser := &models.User{
                PhoneNumber: phoneNumber,
            }
            
            err = s.userRepo.CreateUser(newUser)
            if err != nil {
                return "", err
            }
            user = newUser
        } else {
            return "", err
        }
    }
    
    expiryTime := time.Now().Add(5 * time.Minute)
    authToken := &models.AuthToken{
        UserId:     user.Id,
        Token:      smsToken,
        ExpiryDate: &expiryTime,
    }

    authTokenRepo := repository.NewAuthTokenRepository(s.userRepo.Db)
    if err := authTokenRepo.CreateToken(authToken); err != nil {
        return "", err
    }
    
    err = utils.SendVerificationSMS(phoneNumber, smsToken)
    if err != nil {
        return "", err
    }
    
    return smsToken, nil
}

func (s *UserService) ValidateTokenAndLogin(phoneNumber, smsToken string) (string, error) {
    err := s.ValidateSMSToken(phoneNumber, smsToken)
    if err != nil {
        return "", err
    }

    user, err := s.userRepo.GetUserByPhoneNumber(phoneNumber)
    if err != nil {
        return "", err
    }

    token, err := utils.GenerateJWT(user.Id)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (s *UserService) ValidateSMSToken(phoneNumber, token string) error {
    user, err := s.userRepo.GetUserByPhoneNumber(phoneNumber)
    if err != nil {
        return err
    }

    authTokenRepo := repository.NewAuthTokenRepository(s.userRepo.Db)
    storedToken, err := authTokenRepo.GetTokenByUserID(user.Id)
    if err != nil {
        return err
    }

    if storedToken.Token != token {
        return errors.New("invalid token")
    }

    if storedToken.ExpiryDate.Before(time.Now()) {
        return errors.New("token has expired")
    }

    return nil
}
