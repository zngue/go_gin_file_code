package router
import "github.com/gin-gonic/gin"
func Router(g *gin.RouterGroup)  {
	 // 
	MemberRouter(g) 
	UserInfoRouter(g) 
	UserRouter(g)
	UserCateRouter(g) 
	UserGradeRouter(g)
}
