package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/api/member"
)
func MemberRouter(g *gin.RouterGroup)  {
	api:=g.Group("member")
	{
		api.GET("list",member.List)
		api.POST("del_status",member.DelOrStatus)
		api.GET("detail",member.Detail)
		api.POST("edit",member.Edit)
	}
}