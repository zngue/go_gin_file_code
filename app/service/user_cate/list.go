package user_cate

import (
	"github.com/zngue/go_gin_file_code/app/model/user_cate"
	req "github.com/zngue/go_gin_file_code/app/request/user_cate"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*user_cate.UserCate,err error ) {
	dbConn:=s.UserCate().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&user_cate.UserCate{}).Count(&num).Error
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