package user_grade

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.UserGrade().UpdateStatus(request)
	}else{
		return s.UserGrade().Del(request)
	}
}
