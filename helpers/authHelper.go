package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
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

func HashPassword() {

}
