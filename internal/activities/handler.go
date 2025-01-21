package activities

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service ActivityService
}

func NewHandler(service ActivityService) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) HandleReadActivity(c *gin.Context) {
	activity_id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is not a number",
		})
		return
	}

	// call ActivityService
	activity, err := handler.service.GetActivity(activity_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No activity with id" + c.Param("id"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": activity,
	})
}

func (handler *Handler) HandleCreateActivity(c *gin.Context) {
	body := Activity{}
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	activity, serviceError := handler.service.CreateActivity(&body)

	if serviceError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": serviceError,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"activity": activity,
	})
}
