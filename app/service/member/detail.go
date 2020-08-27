package member
import (
	"github.com/zngue/go_gin_file_code/app/model/member"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *member.Member,err error)  {
	return s.Member().GetOne(request)
}
