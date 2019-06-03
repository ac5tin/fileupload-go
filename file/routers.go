package file

import "github.com/gin-gonic/gin"

// Routes defines all the routes
func Routes(router *gin.RouterGroup) {
	router.GET("/test", Test)
	router.PUT("/upload", Upload)
}
