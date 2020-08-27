package member

import (
	"github.com/zngue/go_gin_file_code/app/model/member"
	req "github.com/zngue/go_gin_file_code/app/request/member"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*member.Member,err error ) {
	dbConn:=s.Member().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&member.Member{}).Count(&num).Error
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