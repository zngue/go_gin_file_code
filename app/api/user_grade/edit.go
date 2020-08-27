package user_grade
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_file_code/app/model/user_grade"
	"github.com/zngue/go_tool/src/common/response"
)
//编辑
func Edit(ctx *gin.Context)  {
	var mb user_grade.UserGrade
	if err:=ctx.ShouldBind(&mb); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if id,err:=Service().Edit(mb); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpFail(ctx,id)
		return
	}
}
