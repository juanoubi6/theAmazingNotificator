package notification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"theAmazingNotificator/app/models"
)

func GetNotifications(c *gin.Context) {

	userID := c.MustGet("id").(uint)
	limit := c.MustGet("limit").(int)
	offset := c.MustGet("offset").(int)

	notificationsData, quantity, err := models.GetAllNotifications(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Something went wrong", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": map[string]interface{}{"notifications": notificationsData, "quantity": quantity}})

}
