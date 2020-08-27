package member

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.Member().UpdateStatus(request)
	}else{
		return s.Member().Del(request)
	}
}
