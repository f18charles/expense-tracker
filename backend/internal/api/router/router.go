package api

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/api/handlers"
	"github.com/f18charles/piggy-bank/backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1
	v1 := r.Group("api/v1")

	authHandler := handlers.NewAuthHandler()
	accountHandler := handlers.NewAccHandler()
	txHandler := handlers.NewTxHandler()

	// public routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	mpesa := v1.Group("/mpesa")
	{
		mpesa.POST("/callback", handlers.MpesaCallback)
	}

	// auth required
	protected := v1.Group("")
	protected.Use(middleware.AuthRequired())
	{
		// auth
		protected.POST("/auth/logout", authHandler.Logout)
		protected.GET("/auth/profile", authHandler.Profile)

		// Accounts
		protected.GET("/accounts", accountHandler.ListAccounts)
		protected.POST("/accounts", accountHandler.CreateAccount)
		protected.GET("/accounts/:id", accountHandler.GetAccount)
		protected.PUT("/accounts/:id", accountHandler.UpdateAccount)
		protected.DELETE("/accounts/:id", accountHandler.DeleteAccount)

		// Transactions
		protected.GET("/transactions", txHandler.ListTransactions)
		protected.POST("/transactions", txHandler.CreateTransactions)
		protected.GET("/transactions/:id", txHandler.GetTransaction)
		protected.PUT("/transactions/:id", txHandler.UpdateTransaction)
		protected.GET("/transactions/export", txHandler.ExportTransactions)

		// Categories
		protected.GET("/categories", handlers.ListCategories)
		protected.POST("/categories", handlers.CreateCategory)
		protected.PUT("/categories/:id", handlers.UpdateCategory)
		protected.DELETE("/categories/:id", handlers.DeleteCategory)

		// Budgets
		protected.GET("/budgets", handlers.Listbudgets)
		protected.POST("/budgets", handlers.CreateBudget)
		protected.GET("/budgets/:id", handlers.GetBudget)
		protected.PUT("/budgets/:id", handlers.UpdateBudget)
		protected.DELETE("/budgets/:id", handlers.DeleteBudget)

		//	Goals
		protected.GET("/goals", handlers.ListGoals)
		protected.POST("/goals", handlers.CreateGoal)
		protected.GET("/goals/:id", handlers.GetGoal)
		protected.PUT("/goals/:id", handlers.UpdateGoal)
		protected.DELETE("/goals/:id", handlers.DeleteGoal)

		// Summary & Insights
		protected.GET("/summary/monthly", handlers.MonthlySummary)
		protected.GET("/summary/overview", handlers.Overview)
		protected.GET("/summary/spending", handlers.SpendingInsights)

		// Mpesa(authenticated)
		protected.POST("/mpesa/stk-push", handlers.MpesaSTKPush)
		protected.GET("/mpesa/status/:id", handlers.MpesaStatus)

		// bank
	}

	return r
}
