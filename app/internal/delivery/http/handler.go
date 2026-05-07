package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(r gin.IRouter) {
	r.GET("/system", h.GetInfo)
}

func (h *Handler) GetInfo(ctx *gin.Context) {
	info := h.service.GetInfo()
	ctx.JSON(http.StatusOK, gin.H{
		"cpu_model":       info.CpuModel,
		"memory_total_kb": info.MemoryTotalKb,
		"hostname":        info.Hostname,
		"memory_limit":    info.MemoryLimit,
	})
}
