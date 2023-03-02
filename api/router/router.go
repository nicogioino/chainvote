package router

import (
	"chain-vote-api/middlewares"
	"chain-vote-api/security"
	"chain-vote-api/services"
	"github.com/gin-gonic/gin"
)

func intializePrivateRouter(engine *gin.Engine) {
	protected := engine.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", services.CurrentUser)
	protected.PUT("/register-address", services.RegisterETHAddress)
	protected.POST("/election", services.CreateElection)
	protected.GET("/election/:id", services.GetElectionById)
	protected.GET("/elections", services.GetAllElections)
}

func initializePublicRouter(engine *gin.Engine) {
	public := engine.Group("/api")
	public.POST("/register", services.Register)
	public.POST("/login", security.Login)
	public.GET("/health", services.HealthCheck)
}

func InitializeRouter(engine *gin.Engine) {
	initializePublicRouter(engine)
	intializePrivateRouter(engine)
}
