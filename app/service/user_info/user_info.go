package user_info
import "github.com/zngue/go_gin_file_code/app/model/user_info"
type Service struct {
}
func (Service)UserInfo() *user_info.UserInfo  {
	return new(user_info.UserInfo)
}
