package {{model}}
import (
	"{{GoModelPath}}/{{modelFile}}"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *{{model}}.{{StructName}},err error)  {
	return s.{{StructName}}().GetOne(request)
}
