package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/api/user"
)
func UserRouter(g *gin.RouterGroup)  {
	api:=g.Group("user")
	{
		api.GET("list",user.List)
		api.POST("del_status",user.DelOrStatus)
		api.GET("detail",user.Detail)
		api.POST("edit",user.Edit)
	}
}