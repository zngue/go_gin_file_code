package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/api/user_grade"
)
func UserGradeRouter(g *gin.RouterGroup)  {
	api:=g.Group("user_grade")
	{
		api.GET("list",user_grade.List)
		api.POST("del_status",user_grade.DelOrStatus)
		api.GET("detail",user_grade.Detail)
		api.POST("edit",user_grade.Edit)
	}
}