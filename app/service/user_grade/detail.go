package user_grade
import (
	"github.com/zngue/go_gin_file_code/app/model/user_grade"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *user_grade.UserGrade,err error)  {
	return s.UserGrade().GetOne(request)
}
