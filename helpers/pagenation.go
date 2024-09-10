package helpers

import (
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
)
func GetPaginationParams(ctx *gin.Context) (int, int, error) {
	page := 1
	limit := 10

	pageParam := ctx.Query("page")
	if pageParam != "" {
		pageNum, err := strconv.Atoi(pageParam)
		if err != nil || pageNum <= 0 {
			return 0, 0, errors.New("유효하지 않은 page 값입니다.")
		}
		page = pageNum
	}

	limitParam := ctx.Query("limit")
	if limitParam != "" {
		limitNum, err := strconv.Atoi(limitParam)
		if err != nil || limitNum <= 0 {
			return 0, 0, errors.New("유효하지 않은 limit 값입니다.")
		}
		limit = limitNum
	}

	return page, limit, nil
}