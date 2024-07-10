package auth

import (
	"context"
	"errors"
	// "strconv"
	"time"

	// "os"
	"log"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/repository"
	"golang.org/x/e-calender/model"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("the secret of kalimdor")
	LOGIN_EXPIRATION_DURATION = 10 * time.Hour
)

var (
	ErrUserNotFound    = errors.New("failed to find user")
	ErrInvalidPassword = errors.New("passwords do not match")
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

// func (u *UserService) UploadImage() {

// }

func NewAuthService(authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}

func (u *AuthService) CreateAccount(ctx context.Context, user *dto.UserDto) error {
	// var userRegistered *model.User
	// phoneNumber, err := strconv.Atoi(user.PhoneNumber)
	// if err != nil {
	// 	return err
	// }

	userRegistered := &entity.UserEntity {
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		CreatedAt: time.Now(),
	}

	err := u.AuthRepository.Create(ctx, userRegistered)
	if err != nil {
		log.Println("user")
		return errors.New("failed to create acc in service layer")
	}

	return nil
}

func (u *AuthService) Find(ctx context.Context, user *dto.UserDto) (string, error) {
	userFetched, err := u.AuthRepository.FindAcc(ctx, user.Username)
	if err != nil {
		log.Println("Failed to find user:", err)
		return "", ErrUserNotFound
	}

	// Compare passwords securely
	if !comparePasswords(userFetched.Password, user.Password) {
		return "", ErrInvalidPassword
	}

	expirationTime := time.Now().Add(LOGIN_EXPIRATION_DURATION)
	expire := jwt.NewNumericDate(expirationTime)

	claims := &model.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
		Username: userFetched.Username,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		log.Println("Failed to sign JWT token:", err)
		return "", errors.New("failed to sign JWT token")
	}

	// Generate and store refresh token
	// refreshToken, err := generateRefreshToken()
	// if err != nil {
	//     return "", errors.New("failed to generate refresh token")
	// }

	// Store refresh token securely

	return signedToken, nil
}

func comparePasswords(hashedPassword, password string) bool {
	// Implement a secure password comparison method, such as bcrypt.CompareHashAndPassword
	// Example:
	// err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// return err == nil
	return hashedPassword == password // Temporary insecure comparison for demonstration
}

func (u *AuthService) ErrUserNotFound() error {
	return ErrUserNotFound
}

func (u *AuthService) ErrInvalidPassword() error {
	return ErrInvalidPassword
}

func (u *AuthService) GetUserInformation(username string) (*dto.UserDto, error) {

	user, err := u.AuthRepository.GetSelfInformation(username)
	if err != nil {
		return nil, err
	}

	data := &dto.UserDto{
		Username: username,
		Email: user.Email,
		// PhoneNumber: user.PhoneNumber,

	}

	return data, nil
}



// func (u *AuthService) Refresh() (string, error) {

// }

// func (u *AuthService) UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	_, err = uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fileName),
// 		Body:   file,
// 	})

// 	return err
// }
