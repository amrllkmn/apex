package activities

import (
	"fmt"
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
			"message": "No activity with id " + c.Param("id"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"activity": activity,
	})
}

func (handler *Handler) HandleCreateActivity(c *gin.Context) {
	body := Activity{}
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, serviceError := handler.service.CreateActivity(&body)
	fmt.Println(body)
	if serviceError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": serviceError.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"activity": body,
	})
}

func (handler *Handler) HandleGetActivities(c *gin.Context) {
	activities, err := handler.service.GetActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"activities": activities,
	})
}

func (handler *Handler) HandleUpdateActivity(c *gin.Context) {
	activity_id, convErr := strconv.Atoi(c.Param("id"))

	if convErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is not a number",
		})
		return
	}

	body := Activity{}
	bindErr := c.BindJSON(&body)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": bindErr.Error(),
		})
		return
	}

	svcMsg, svcErr := handler.service.UpdateActivity(activity_id, &body)
	if svcErr != nil {
		fmt.Println(svcErr)
		c.JSON(http.StatusNotFound, gin.H{
			"message": svcMsg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": svcMsg,
	})
}

func (handler *Handler) HandleDeleteActivity(c *gin.Context) {
	activity_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is not a number",
		})
		return
	}

	msg, err := handler.service.DeleteActivity(activity_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No activity with id" + c.Param("id"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
