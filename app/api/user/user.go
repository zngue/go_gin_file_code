package user
import "github.com/zngue/go_gin_file_code/app/service/user"
func Service() *user.Service  {
	return new(user.Service)
}
