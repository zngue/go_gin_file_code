package member

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Member struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	db.TimeStampModel
}

func (m *Member) TableName() string  {
	return "member"
}
//添加
func (m *Member) Add(Member Member) (ID int,err error) {
	err=db.MysqlConn.Create(&Member).Error
	ID = Member.ID
	return
}
//修改
func (m *Member) Update(data Member ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Member{}).Update(&dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Member) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Member{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Member) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Member{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Member{}).Error
	return
}
func (m *Member) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Member) GetOne (request request.IDStatusRequest ) ( mb *Member,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






