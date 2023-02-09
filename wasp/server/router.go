package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/berachain/stargazer/wasp/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(logger *zap.Logger, handler *handler.Handler) *gin.Engine {
	fmt.Println("Initializing router...")
	engine := gin.Default()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(gin.Recovery())
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.GET("/block", handler.GetLatestBlock)
	engine.GET("/block/:height", handler.GetBlock)
	engine.GET("/blocks", handler.GetBlocks)
	return engine
}
