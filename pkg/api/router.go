package api

import (
	_ "owwi/docs"
	"owwi/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	// This function should return a new Gin engine with routes defined.
	// For now, we will return a new Gin engine without any routes.
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// router.GET("/health", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"status":  "ok",
	// 		"message": "Service is running",
	// 	})
	// })
	router.Use(middleware.CORSMiddleware())

	// Define authentication routes
	router.POST("/login", AuthApi.Login)
	router.POST("/register", AuthApi.Register)
	router.GET("/whoami", middleware.IsUser, AuthApi.WhoAmI)

	// Define routes for types
	router.POST("/types", middleware.IsUser, TypeApi.CreateType)
	router.GET("/types", middleware.IsUser, TypeApi.GetTypes)

	// Define routes for categories
	categoryGroup := router.Group("/categories", middleware.IsUser)
	categoryGroup.POST("", CategoryApi.CreateCategory)
	categoryGroup.PUT("/:id", CategoryApi.UpdateCategory)
	categoryGroup.GET("", CategoryApi.GetAllCategoriesByUser)
	categoryGroup.GET("/:id", CategoryApi.GetCategoryByID)
	categoryGroup.DELETE("/:id", CategoryApi.DeleteCategory)

	// Define routes for partners
	partnerGroup := router.Group("/partners", middleware.IsUser)
	partnerGroup.POST("", PartnerApi.CreatePartner)
	partnerGroup.PUT("/:id", PartnerApi.UpdatePartner)
	partnerGroup.GET("", PartnerApi.GetAllPartnersByUser)
	partnerGroup.GET("/:id", PartnerApi.GetPartnerByID)
	partnerGroup.DELETE("/:id", PartnerApi.DeletePartner)


	// Define routes for transactions
	router.POST("/transactions", middleware.IsUser, TransactionApi.CreateTransaction)
	// router.PUT("/transactions/:id", middleware.IsUser, TransactionApi.UpdateTransaction)
	// router.GET("/transactions", middleware.IsUser, TransactionApi.GetAllTransactionsByUser)
	// router.GET("/transactions/:id", middleware.IsUser, TransactionApi.GetTransactionByID)
	// router.DELETE("/transactions/:id", middleware.IsUser, TransactionApi.DeleteTransaction)

	// // Define routes for reports
	// router.GET("/reports/transactions/weekly", middleware.IsUser, ReportApi.GetTransactionsReport)
	// router.GET("/reports/transactions/monthly", middleware.IsUser, ReportApi.GetCategoriesReport)

	/* ---------- INCOMING FEATURES ---------- */
	// Define routes for users
	// router.GET("/users", middleware.IsAdmin, UserApi.GetAllUsers)
	// router.GET("/users/:id", middleware.IsAdmin, UserApi.GetUserByID)
	// router.PUT("/users/:id", middleware.IsAdmin, UserApi.UpdateUser)
	// router.DELETE("/users/:id", middleware.IsAdmin, UserApi.DeleteUser)
	// // Define routes for user profile
	// router.GET("/profile", middleware.IsUser, UserApi.GetProfile)
	// router.PUT("/profile", middleware.IsUser, UserApi.UpdateProfile)
	// // Define routes for user password
	// router.PUT("/profile/password", middleware.IsUser, UserApi.UpdatePassword)
	// // Define routes for user avatar
	// router.PUT("/profile/avatar", middleware.IsUser, UserApi.UpdateAvatar)
	// // Define routes for user preferences
	// router.GET("/profile/preferences", middleware.IsUser, UserApi.GetPreferences)
	// router.PUT("/profile/preferences", middleware.IsUser, UserApi.UpdatePreferences)
	// // Define routes for user notifications
	// router.GET("/profile/notifications", middleware.IsUser, UserApi.GetNotifications)
	// router.PUT("/profile/notifications", middleware.IsUser, UserApi.UpdateNotifications)
	// // Define routes for user security
	// router.GET("/profile/security", middleware.IsUser, UserApi.GetSecuritySettings)
	// router.PUT("/profile/security", middleware.IsUser, UserApi.UpdateSecuritySettings)
	// // Define routes for user activity logs
	// router.GET("/profile/activity-logs", middleware.IsUser, UserApi.GetActivityLogs)
	// // Define routes for user sessions
	// router.GET("/profile/sessions", middleware.IsUser, UserApi.GetSessions)
	// router.DELETE("/profile/sessions/:id", middleware.IsUser, UserApi.DeleteSession)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
