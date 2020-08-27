package {{model}}
import "{{GoModelPath}}/{{modelFile}}"
func (s *Service) Edit(c {{model}}.{{StructName}}) (int ,error) {
	if c.ID>0 {
		return s.{{StructName}}().Update(c)
	}else {
		return s.{{StructName}}().Add(c)
	}
}
