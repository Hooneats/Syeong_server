package router

import (
	"github.com/Hooneats/Syeong_server/logger"
	v1 "github.com/Hooneats/Syeong_server/router/handler/v1"
	"github.com/Hooneats/Syeong_server/router/middleware"
	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
)

var ginR *GinRoute

type GinRoute struct {
	engin *gin.Engine
}

func NewGinRoute(mode string) *GinRoute {
	if ginR != nil {
		return ginR
	}
	setMode(mode)
	ginR = &GinRoute{
		engin: newEngine(),
	}
	return ginR
}

func (r *GinRoute) Handle() http.Handler {
	gr := r.engin

	registerSwagger(gr)

	v1.GinHandle(gr)

	return r.engin
}

// newEngine generate gin engin and global middleware setting
func newEngine() *gin.Engine {
	grt := gin.New()
	setMiddleware(grt)
	return grt
}

func setMode(mode string) {
	switch mode {
	case "dev":
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	case "prod":
		logger.AppLog.Info("Start gin mod", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	case "test":
		logger.AppLog.Info("Start gin mod", gin.TestMode)
		gin.SetMode(gin.TestMode)
	default:
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	}
}

func setMiddleware(grt *gin.Engine) {
	grt.Use(middleware.GinLogger())
	grt.Use(middleware.GinRecovery(true))
	grt.Use(middleware.CORS())
}

func registerSwagger(gr *gin.Engine) {
	// TODO Now, Swagger is Default Swagger. Should change this to custom swagger
	gr.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
}
