package {{model}}
import "{{GoModelPath}}/{{serviceFile}}"
func Service() *{{model}}.Service  {
	return new({{model}}.Service)
}
