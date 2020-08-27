package user

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type User struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *User) TableName() string  {
	return "user"
}
//添加
func (m *User) Add(User User) (ID int,err error) {
	err=db.MysqlConn.Create(&User).Error
	ID = User.ID
	return
}
//修改
func (m *User) Update(data User ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&User{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *User) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&User{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *User) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&User{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&User{}).Error
	return
}
func (m *User) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *User) GetOne (request request.IDStatusRequest ) ( mb *User,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






