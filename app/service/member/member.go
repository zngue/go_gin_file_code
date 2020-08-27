package member
import "github.com/zngue/go_gin_file_code/app/model/member"
type Service struct {
}
func (Service)Member() *member.Member  {
	return new(member.Member)
}
