package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/web-cert-checker/api/v1/handlers"
)

//SetupRouter sets up the routes for the reuters api
func SetupRouter() *gin.Engine {
	//get GIN_MODE variable
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/test", handlers.TestHandler)
	v1.POST("/sslchecker", handlers.SSLCheckerHandler)

	return router
}
