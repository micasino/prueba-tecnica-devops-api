package helper

import (
	"mg-inc-devops/pkg/domain"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) domain.Pagination {
	limit := 1
	page := 1
	//sort := `created_at DESC`
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "per_page":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
			// case "sort":
			// 	sort = queryValue
			// 	break
		}
	}

	return domain.Pagination{
		Limit: limit,
		Page:  page,
		//Sort:  sort,
	}
}

func Paginate(pagination *domain.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit := pagination.Limit
		page := pagination.Page

		offset := (page - 1) * limit

		return db.Offset(int(offset)).Limit(int(limit))
	}
}

// func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		limit, _ := c.MustGet(constants.Limit).(int64)
// 		page, _ := c.MustGet(constants.Page).(int64)

// 		offset := (page - 1) * limit

// 		return db.Offset(int(offset)).Limit(int(limit))
// 	}
// }

// func paginate(value interface{}, pagination *domain.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
// 	var totalRows int64
// 	db.Model(value).Count(&totalRows)
// 	pagination.TotalRows = totalRows
// 	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
// 	pagination.TotalPages = totalPages

// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()) //.Order(pagination.GetSort())
// 	}
// }

// func GeneratePaginationFromRequest(c *gin.Context) domain.Pagination {
// 	// Initializing default
// 	//	var mode string
// 	limit := 2
// 	page := 1
// 	sort := "created_at asc"
// 	query := c.Request.URL.Query()
// 	for key, value := range query {
// 		queryValue := value[len(value)-1]
// 		switch key {
// 		case "limit":
// 			limit, _ = strconv.Atoi(queryValue)
// 			break
// 		case "page":
// 			page, _ = strconv.Atoi(queryValue)
// 			break
// 		case "sort":
// 			sort = queryValue
// 			break

// 		}
// 	}
// 	return domain.Pagination{
// 		Limit: limit,
// 		Page:  page,
// 		Sort:  sort,
// 	}

// }
