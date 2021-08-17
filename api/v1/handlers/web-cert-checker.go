package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/web-cert-checker/api/v1/controllers"
)

//TestHandler a test handler
func TestHandler(ctx *gin.Context) {
	responce := map[string]string{
		"status":  "success",
		"message": "Hello world",
	}
	ctx.JSON(http.StatusOK, responce)
}

//SSLCheckerHandler function that check SSL certificate
func SSLCheckerHandler(ctx *gin.Context) {
	// Get host and port from request
	host := ctx.Request.FormValue("host")
	if host == "" {
		responce := map[string]string{
			"status":  "error",
			"message": "you must provide a hostname",
		}
		ctx.JSON(http.StatusBadRequest, responce)
		return
	}
	certificateData, err := controllers.SSLChecker(host)
	if err != nil {
		responce := map[string]string{
			"status":  "error",
			"message": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, responce)
		return
	}

	ctx.JSON(http.StatusOK, certificateData)
}
