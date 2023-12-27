package middleware

import (
	"mg-inc-devops/pkg/constants"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(c *gin.Context) {
	perPage, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
	if err != nil {
		perPage = 10
	}

	page, err := strconv.ParseInt(c.Query("page"), 10, 0)
	if err != nil {
		page = 0
	}

	c.Set(constants.Limit, perPage)
	c.Set(constants.Page, page)

	c.Next()
}
