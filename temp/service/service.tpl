package {{model}}
import "{{GoModelPath}}/{{modelFile}}"
type Service struct {
}
func (Service){{StructName}}() *{{model}}.{{StructName}}  {
	return new({{model}}.{{StructName}})
}
