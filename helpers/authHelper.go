package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("unauthorized for this area")
		return err
	}
	return err
}
func MathUserTypeToUId(c *gin.Context, id string) (err error) {
	userType := c.GetString("user_type")
	uId := c.GetString("uid")
	err = nil

	if userType == "USER" && uId != id {
		err = errors.New("unauthorized for this area")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func VerifyPassword(providedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	mess := ""
	correct := true

	if err != nil {
		mess = "password is wrong or incorrect"
		correct = false
		return correct, errors.New(mess)

	}
	return correct, nil

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)

}
