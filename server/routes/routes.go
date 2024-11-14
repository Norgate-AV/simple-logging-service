package routes

import (
	"github.com/gin-gonic/gin"
)

const (
	BaseRouteV1 = "/api/v1"
)

func Register(server *gin.Engine) {
	server.GET(BaseRouteV1+"/health", GetHealth)
	server.POST(BaseRouteV1+"/log", CreateLog)
}
