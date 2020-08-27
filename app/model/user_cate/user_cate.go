package user_cate

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type UserCate struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *UserCate) TableName() string  {
	return "user_cate"
}
//添加
func (m *UserCate) Add(UserCate UserCate) (ID int,err error) {
	err=db.MysqlConn.Create(&UserCate).Error
	ID = UserCate.ID
	return
}
//修改
func (m *UserCate) Update(data UserCate ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&UserCate{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *UserCate) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&UserCate{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *UserCate) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&UserCate{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&UserCate{}).Error
	return
}
func (m *UserCate) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *UserCate) GetOne (request request.IDStatusRequest ) ( mb *UserCate,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






