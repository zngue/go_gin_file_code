package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/api/user_cate"
)
func UserCateRouter(g *gin.RouterGroup)  {
	api:=g.Group("user_cate")
	{
		api.GET("list",user_cate.List)
		api.POST("del_status",user_cate.DelOrStatus)
		api.GET("detail",user_cate.Detail)
		api.POST("edit",user_cate.Edit)
	}
}