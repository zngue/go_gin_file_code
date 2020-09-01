package {{model}}
import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
)
type ListRequest struct {
	IDString string `json:"id_string" form:"id_string"`
	IDArr []int
	DBConn DBConn
	Order  DBConn
	request.Page
}
type DBConn func(db *gorm.DB) *gorm.DB
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
	l.Order= func(db *gorm.DB) *gorm.DB {
		return db
	}
}
