package router
import (
	"github.com/gin-gonic/gin"
	"{{GoModelPath}}/{{apiFile}}"
)
func {{StructName}}Router(g *gin.RouterGroup)  {
	api:=g.Group("{{model}}")
	{
		api.GET("list",{{model}}.List)
		api.POST("del_status",{{model}}.DelOrStatus)
		api.GET("detail",{{model}}.Detail)
		api.POST("edit",{{model}}.Edit)
	}
}