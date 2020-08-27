package user_info
import (
	"github.com/zngue/go_gin_file_code/app/model/user_info"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *user_info.UserInfo,err error)  {
	return s.UserInfo().GetOne(request)
}
