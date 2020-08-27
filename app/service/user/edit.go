package user
import "github.com/zngue/go_gin_file_code/app/model/user"
func (s *Service) Edit(c user.User) (int ,error) {
	if c.ID>0 {
		return s.User().Update(c)
	}else {
		return s.User().Add(c)
	}
}
