package user_cate
import (
	"github.com/zngue/go_gin_file_code/app/model/user_cate"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *user_cate.UserCate,err error)  {
	return s.UserCate().GetOne(request)
}
