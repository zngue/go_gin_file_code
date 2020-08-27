package {{model}}
import (
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/fun/zng_str"
)
type ListRequest struct {
	IDString string `json:"id_string" form:"id_string"`
	IDArr []int
	request.Page
}
func (l *ListRequest) TypeStringToArr()  {
	l.IDArr = zng_str.IDStringToSlice(l.IDString,",")
}
