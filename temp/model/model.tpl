package {{model}}

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type {{modelName}} struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *{{modelName}}) TableName() string  {
	return "{{model}}"
}
//添加
func (m *{{modelName}}) Add({{modelName}} {{modelName}}) (ID int,err error) {
	err=db.MysqlConn.Create(&{{modelName}}).Error
	ID = {{modelName}}.ID
	return
}
//修改
func (m *{{modelName}}) Update(data {{modelName}} ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&{{modelName}}{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *{{modelName}}) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&{{modelName}}{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *{{modelName}}) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&{{modelName}}{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&{{modelName}}{}).Error
	return
}
func (m *{{modelName}}) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *{{modelName}}) GetOne (request request.IDStatusRequest ) ( mb *{{modelName}},err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






