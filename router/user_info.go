package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/api/user_info"
)
func UserInfoRouter(g *gin.RouterGroup)  {
	api:=g.Group("user_info")
	{
		api.GET("list",user_info.List)
		api.POST("del_status",user_info.DelOrStatus)
		api.GET("detail",user_info.Detail)
		api.POST("edit",user_info.Edit)
	}
}