package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Creates a response with given message and data
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{"Status": true, "Data": data})
}

// Aborts request with an error response with given message
func FailResponse(c *gin.Context, code int) {
	c.AbortWithStatusJSON(code, map[string]interface{}{"Status": false, "Data": []interface{}{}})
}
