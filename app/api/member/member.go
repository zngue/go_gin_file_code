package member
import "github.com/zngue/go_gin_file_code/app/service/member"
func Service() *member.Service  {
	return new(member.Service)
}
