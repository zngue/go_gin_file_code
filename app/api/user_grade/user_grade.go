package user_grade
import "github.com/zngue/go_gin_file_code/app/service/user_grade"
func Service() *user_grade.Service  {
	return new(user_grade.Service)
}
