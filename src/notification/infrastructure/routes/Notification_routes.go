package routes

import (
	"PubNotification/src/notification/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesAsignature(
	r *gin.Engine,
	createAsignatureController *controllers.CreateAsignatureController,

) {
	r.POST("/send-notification", createAsignatureController.Execute)

}
