package responses

import (
	"mg-inc-devops/pkg/constants"

	"github.com/gin-gonic/gin"
)

// JSON : json response function
func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": "error",
		"code":    statusCode,
		"data":    data,
		"status":  false,
	})
}

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": "success",
		"code":    statusCode,
		"data":    data,
		"status":  true,
	})
}

// JSONWithPagination : json response function
func JSONWithPagination(c *gin.Context, statusCode int, data interface{}, count int64) {
	limit, _ := c.MustGet(constants.Limit).(int64)
	size, _ := c.MustGet(constants.Page).(int64)
	// limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	// size, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	c.JSON(
		200,
		gin.H{
			"message":    "success",
			"code":       statusCode,
			"data":       data,
			"pagination": gin.H{"has_next": (count - limit*size) > 0, "count": count},
			"status":     true,
		})
}
