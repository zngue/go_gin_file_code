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
func (l *ListRequest) ListInit()  {
	l.ListDbConn()
	l.ListOrder()
}
func (l *ListRequest) ListDbConn() {
	l.DBConn= func(db *gorm.DB) *gorm.DB {
		return db
	}
}
func (l *ListRequest) ListOrder()  {
	l.DBConn= func(db *gorm.DB) *gorm.DB {
		return db
	}
}

