package user_info

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type UserInfo struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *UserInfo) TableName() string  {
	return "user_info"
}
//添加
func (m *UserInfo) Add(UserInfo UserInfo) (ID int,err error) {
	err=db.MysqlConn.Create(&UserInfo).Error
	ID = UserInfo.ID
	return
}
//修改
func (m *UserInfo) Update(data UserInfo ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&UserInfo{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *UserInfo) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&UserInfo{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *UserInfo) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&UserInfo{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&UserInfo{}).Error
	return
}
func (m *UserInfo) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *UserInfo) GetOne (request request.IDStatusRequest ) ( mb *UserInfo,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






