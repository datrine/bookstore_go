package middleware

import (
	"strings"

	"github.com/datrine/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader, ok := c.Request.Header["Authorization"]
	if !ok {
		c.AbortWithStatusJSON(401, &utils.ErrResponse{
			Message: "Unauthorized stuff",
		})
	}
	token := strings.Split(authHeader[0], " ")[1]
	if token == "" {
		c.AbortWithStatusJSON(401, &utils.ErrResponse{
			Message: "Unauthorized",
		})
		return
	}
	payload, err := utils.VerifyAccessToken(token)
	if err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(401, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	c.Keys = make(map[string]interface{})
	if c.Keys == nil {
		c.AbortWithStatusJSON(401, &utils.ErrResponse{
			Message: "Nil map",
		})
		return
	}
	c.Set("user", payload)
}
