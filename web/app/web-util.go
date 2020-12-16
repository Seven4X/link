package app

import (
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetUserId(e echo.Context) int {
	u := e.Get(consts.User)
	if u == nil {
		return 0
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	return claims.Id
}