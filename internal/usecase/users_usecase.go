package usecase

import (
	"context"
	"errors"
	"golang-redis/internal/entity"
	"golang-redis/internal/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// type UsersUseCase interface {
// 	Login(ctx context.Context, email string) (string, *entity.MasterUsers, error)
// 	Register(ctx context.Context, username, email string) (*entity.MasterUsers, error)
// 	ValidateJWT(ctx context.Context, tokenString string) (*entity.MasterUsers, error)
// }

type UsersUseCase struct {
	Repository *repository.UsersRepository
	JWTKey     []byte
}

// NewUsersUseCase constructor
func NewUsersUseCase(repository *repository.UsersRepository) *UsersUseCase {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	return &UsersUseCase{
		Repository: repository,
		JWTKey:     jwtKey,
	}
}

func (uc *UsersUseCase) ValidateJWT(ctx context.Context, tokenString string) (*entity.MasterUsers, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return uc.JWTKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid token payload")
	}

	user, err := uc.Repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc *UsersUseCase) Login(ctx context.Context, email string) (string, *entity.MasterUsers, error) {
	user, err := uc.Repository.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, errors.New("users not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(uc.JWTKey)
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (uc *UsersUseCase) Register(ctx context.Context, username string, email string) (*entity.MasterUsers, error) {
	// validate email
	existing, _ := uc.Repository.GetByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("EMAIL_ALREADY_REGISTERED")
	}

	// create users if not exist
	users, err := uc.Repository.Create(ctx, username, email)
	if err != nil {
		return nil, err
	}

	return users, nil
}
