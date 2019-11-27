package browser

import "github.com/gin-gonic/gin"

// Routes ...
func Routes(router *gin.Engine) {
	browserAPI := router.Group("/browser")
	{
		browserAPI.GET("/launch", Launch)
	}
}
