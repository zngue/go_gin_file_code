package user_info
import "github.com/zngue/go_gin_file_code/app/service/user_info"
func Service() *user_info.Service  {
	return new(user_info.Service)
}
