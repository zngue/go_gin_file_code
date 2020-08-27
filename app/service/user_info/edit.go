package user_info
import "github.com/zngue/go_gin_file_code/app/model/user_info"
func (s *Service) Edit(c user_info.UserInfo) (int ,error) {
	if c.ID>0 {
		return s.UserInfo().Update(c)
	}else {
		return s.UserInfo().Add(c)
	}
}
