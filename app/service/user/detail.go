package user
import (
	"github.com/zngue/go_gin_file_code/app/model/user"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *user.User,err error)  {
	return s.User().GetOne(request)
}
