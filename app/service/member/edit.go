package member
import "github.com/zngue/go_gin_file_code/app/model/member"
func (s *Service) Edit(c member.Member) (int ,error) {
	if c.ID>0 {
		return s.Member().Update(c)
	}else {
		return s.Member().Add(c)
	}
}
