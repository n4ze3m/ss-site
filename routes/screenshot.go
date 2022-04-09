package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/n4ze3m/ss-site/controllers"
)

func Screenshot(r *gin.Engine) {

	r.GET("/screenshot", controllers.GenerateScreenshot)
}
