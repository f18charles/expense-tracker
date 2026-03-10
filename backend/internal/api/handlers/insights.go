package handlers

import (
	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/pkg/summary"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SummaryHandler struct {
	summaryService *services.SummaryService
}

func NewSummaryHandler(db *gorm.DB) *SummaryHandler {
	return &SummaryHandler{
		summaryService: services.NewSummaryService(db),
	}
}

// MonthlySummary returns aggregated monthly summary data for the user.
func (sh *SummaryHandler) MonthlySummary(c *gin.Context) {

}

// Overview returns a high-level overview/dashboard for the user.
func Overview(c *gin.Context) {}

// SpendingInsights returns insights and breakdowns for spending patterns.
func SpendingInsights(c *gin.Context) {}
