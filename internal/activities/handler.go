package activities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello(c *gin.Context) {
	name := c.Param("name")
	message := "Hello " + name
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
