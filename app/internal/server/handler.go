package server

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(r gin.IRouter)
	GetInfo(ctx *gin.Context)
}
