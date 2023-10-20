package routes

import (
	"github.com/datrine/repositories"
	"github.com/datrine/services"
	"github.com/datrine/utils"
	"github.com/gin-gonic/gin"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.OkResponse
// @Failure 400 {object} utils.ErrResponse
// @Router /auth/login/basic [post]
func basicLogin(c *gin.Context) {
	var b utils.BasicLoginDto
	err := c.Bind(&b)
	if err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(400, &utils.ErrResponse{Message: err.Error()})
		return
	}
	tokens, err := services.BasicLoginService(&b)
	if err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(400, &utils.ErrResponse{Message: err.Error()})
		return
	}
	c.JSON(200, &utils.OkResponse{Message: "Log in successfull",
		Data: &LoginResponse{AccessToken: tokens.AccessToken}})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description Register user
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.OkResponse
// @Failure 400 {object}  utils.ErrResponse
// @Router /auth/register [post]
func registerRoute(c *gin.Context) {
	data := &utils.RegisterDto{}
	err := c.Bind(data)
	if err != nil {
		c.JSON(400, &utils.ErrResponse{Message: err.Error()})
		return
	}
	user, err := repositories.CreateUser(data)
	if err != nil {
		c.JSON(400, &utils.ErrResponse{Message: err.Error()})
		return
	}
	c.JSON(200, user)
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
