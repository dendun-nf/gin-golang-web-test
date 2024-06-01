package commons

import "github.com/gin-gonic/gin"

type IEndpoint interface {
	MapEndpoint(router *gin.RouterGroup)
}
