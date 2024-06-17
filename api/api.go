package api

import (
	"github.com/Salikhov079/library-auth/api/handler"
	"github.com/Salikhov079/library-auth/api/middleware"

	_ "github.com/Salikhov079/library-auth/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Library service
// @version 1.0
// @description Library service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.MiddleWare())
	u := r.Group("/user")
	u.POST("/register", h.RegisterUser)
	u.POST("/login", h.LoginUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/get/:id", h.GetbyIdUser)
	u.GET("/:id/borrowed_books", h.GetBorrowedBooks)
	u.GET("/:id/borrowing_history", h.GetBorrowingHistory)
	
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	return r
}
