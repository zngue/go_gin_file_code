package user_grade
import "github.com/zngue/go_gin_file_code/app/model/user_grade"
func (s *Service) Edit(c user_grade.UserGrade) (int ,error) {
	if c.ID>0 {
		return s.UserGrade().Update(c)
	}else {
		return s.UserGrade().Add(c)
	}
}
