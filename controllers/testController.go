package controllers

import (
	"fmt"
	"net/http"
	"os"

	models "test-echo/model"
	"test-echo/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetTest(c echo.Context) error {
	var Res models.ResponseView

	Res.ResponseCode = http.StatusOK
	Res.ResponseDesc = "SUCESS"
	Res.ResponseTime = utils.DateToStdNow()
	Res.Message = "TEST"
	Res.Result = "TEST"

	return c.JSON(http.StatusOK, Res)

}

func GenerateToken(c echo.Context) error {

	token, errToken := GetTokenJWT("test@mail.com")

	if errToken != nil {
		fmt.Println("Token Error.")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "gagal"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success", "token": token})

}

func GetTokenJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["level"] = "application"
	//claims["exp"] = time.Now().Add(time.Hour * 14).Unix()

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tokens.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		errMsg := "Token Error"
		fmt.Println("Token error.")
		return errMsg, err
	}

	return t, nil
}
