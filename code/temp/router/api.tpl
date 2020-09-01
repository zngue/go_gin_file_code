package router
import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "{{GoModelPath}}/docs"
)
func Router(g *gin.RouterGroup)  {
	//
    g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
