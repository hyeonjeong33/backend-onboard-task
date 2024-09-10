package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다. 입력한 데이터를 확인해주세요."})
		return false
	}
	return true
}