package services

import (
    "errors"
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
    "backend-onboard-task/helpers"
    "backend-onboard-task/models"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) RegisterUser(input models.User) error {
    var existingUser models.User
	err := s.DB.Where("email = ?", input.Email).First(&existingUser).Error
	if err == nil {
		return errors.New("이미 사용 중인 이메일입니다.")
	}

    if len(input.Role) == 0 {
		input.Role = "User"
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return errors.New("비밀번호를 암호화하는 중 오류가 발생했습니다.")
    }
    input.Password = string(hashedPassword)

    if err := s.DB.Create(&input).Error; err != nil {
		return errors.New("회원가입 처리 중 오류가 발생했습니다. 잠시 후 다시 시도해 주세요.")
	}

    return nil
}

func (s *UserService) LoginUser(input models.User) (string, error) {
    var user models.User
    if err := s.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        return "", errors.New("이메일 또는 비밀번호가 일치하지 않습니다. 로그인 정보를 다시 확인해주세요")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return "", errors.New("이메일 또는 비밀번호가 일치하지 않습니다. 로그인 정보를 다시 확인해주세요")
    }

    token, err := helpers.GenerateJWT(user.ID, user.Email)
    if err != nil {
        return "", errors.New("토큰 생성에 실패했습니다.")
    }

    return token, nil
}