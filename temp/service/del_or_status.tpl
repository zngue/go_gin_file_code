package {{model}}

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.{{StructName}}().UpdateStatus(request)
	}else{
		return s.{{StructName}}().Del(request)
	}
}
