package user_grade

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type UserGrade struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *UserGrade) TableName() string  {
	return "user_grade"
}
//添加
func (m *UserGrade) Add(UserGrade UserGrade) (ID int,err error) {
	err=db.MysqlConn.Create(&UserGrade).Error
	ID = UserGrade.ID
	return
}
//修改
func (m *UserGrade) Update(data UserGrade ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&UserGrade{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *UserGrade) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&UserGrade{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *UserGrade) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&UserGrade{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&UserGrade{}).Error
	return
}
func (m *UserGrade) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *UserGrade) GetOne (request request.IDStatusRequest ) ( mb *UserGrade,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






