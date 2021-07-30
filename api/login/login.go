package login

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"fiber-web/api"
	"fiber-web/model"
	"time"
)

// 登录
// @Summary 登录
// @Description 登录
// @Tags login
// @Accept json
// @Produce json
// @Param user body model.User true "login user"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Failure 400 {object} api.ResponseHTTP{}
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return api.Response(c, err, nil)
	}
	err := user.Login()
	if  err != nil{
		return api.Response(c, err, nil)
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, nil, t)
}