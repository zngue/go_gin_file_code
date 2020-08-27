package user_cate
import "github.com/zngue/go_gin_file_code/app/model/user_cate"
func (s *Service) Edit(c user_cate.UserCate) (int ,error) {
	if c.ID>0 {
		return s.UserCate().Update(c)
	}else {
		return s.UserCate().Add(c)
	}
}
