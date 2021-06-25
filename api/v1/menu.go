package v1

import (
	"github.com/gin-gonic/gin"
	"wm-infoflow-api-go/common/response"
)

func MenuList(c *gin.Context) {
	 response.New(c).Format()
}