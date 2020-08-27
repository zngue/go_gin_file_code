package user_grade
import "github.com/zngue/go_gin_file_code/app/model/user_grade"
type Service struct {
}
func (Service)UserGrade() *user_grade.UserGrade  {
	return new(user_grade.UserGrade)
}
