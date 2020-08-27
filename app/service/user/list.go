package user

import (
	"github.com/zngue/go_gin_file_code/app/model/user"
	req "github.com/zngue/go_gin_file_code/app/request/user"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*user.User,err error ) {
	dbConn:=s.User().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&user.User{}).Count(&num).Error
	})
	if b || err!=nil{
		return
	}
	if request.PageSize>0 && request.Page.Page>0 {
		dbConn = dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	dbConn = dbConn.Order("id desc")
	err =dbConn.Find(&mb).Error
	if err != nil {
		return
	}
	return
}