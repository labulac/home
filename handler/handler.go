package handler

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func Gettel(c *gin.Context) {
	tel := viper.GetString("movecar.tel")
	c.String(http.StatusOK, tel)
}

func Changetel(c *gin.Context) {
	tel := c.Param("tel")
	viper.Set("movecar.tel", tel)
	c.String(http.StatusOK, tel)

}
