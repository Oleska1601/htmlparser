package controller

import (
	_ "htmlparser/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) setUpRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/data", s.GetDataHander)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.Server.Handler = router
}
