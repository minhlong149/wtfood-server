package services

import (
	"database/sql"
	"log"
	"os"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/minhlong149/what-the-food/models"
)

type UserService struct {
	Db *sql.DB
}

const (
	USERNAME_PATTERN = "^[a-zA-Z0-9_-]{3,16}$"
	PASSWORD_PATTERN = "^[a-zA-Z0-9_-]{6,18}$"
)

func (s *UserService) AccountIsValid(account models.Account) error {
	if match, err := regexp.MatchString(USERNAME_PATTERN, account.Username); err != nil {
		log.Println(err)
		return err
	} else if !match {
		return models.ErrInvalidUsername
	}

	if match, err := regexp.MatchString(PASSWORD_PATTERN, account.Password); err != nil {
		log.Println(err)
		return err
	} else if !match {
		return models.ErrInvalidPassword
	}

	return nil

}

func (s *UserService) Register(account models.Account) (user models.User, err error) {
	row := s.Db.QueryRow("SELECT username FROM users WHERE username = $1", account.Username)
	err = row.Scan()
	switch err {
	case sql.ErrNoRows:
		user.Username = account.Username
	case nil:
		return models.User{}, models.ErrUsernameExists
	default:
		log.Println(err)
		return models.User{}, err
	}

	hashedPassword, err := hashPassword(account.Password)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	row = s.Db.QueryRow(`
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	`, account.Username, hashedPassword)

	err = row.Scan(&user.Id)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	claims := map[string]interface{}{
		"id":       user.Id,
		"username": user.Username,
	}

	user.Token, err = createToken(claims)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func createToken(claims map[string]interface{}) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
