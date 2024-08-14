package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatusController define o controlador para as rotas de status
type StatusController struct {
	startTime    time.Time
	requestCount int
}

// NewStatusController cria uma nova instância do StatusController
func NewStatusController() *StatusController {
	return &StatusController{
		startTime:    time.Now(),
		requestCount: 0,
	}
}

// RegisterRoutes registra as rotas de status no gin router
func (s *StatusController) RegisterRoutes(r *gin.Engine) {
	// Registro do middleware e da rota
	statusGroup := r.Group("/status")
	statusGroup.Use(s.CountRequests)
	statusGroup.GET("", s.GetStatus)
}

// GetStatus retorna as informações de uptime e contagem de requisições
func (s *StatusController) GetStatus(ctx *gin.Context) {
	upTime := time.Since(s.startTime).String()
	ctx.JSON(http.StatusOK, gin.H{
		"uptime":   upTime,
		"requests": s.requestCount,
	})
}

// CountRequests é um middleware que conta o número de requisições
func (s *StatusController) CountRequests(ctx *gin.Context) {
	s.requestCount++
	ctx.Next()
}

// SetStartTime define o tempo de início
func (s *StatusController) SetStartTime(startTime time.Time) {
	s.startTime = startTime
}

// SetRequestsCount define a contagem de requisições
func (s *StatusController) SetRequestsCount(count int) {
	s.requestCount = count
}
