package user
import "github.com/zngue/go_gin_file_code/app/model/user"
type Service struct {
}
func (Service)User() *user.User  {
	return new(user.User)
}
