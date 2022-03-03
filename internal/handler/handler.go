package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	InitRoutes() *gin.Engine
}
