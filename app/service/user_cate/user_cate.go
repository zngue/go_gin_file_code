package user_cate
import "github.com/zngue/go_gin_file_code/app/model/user_cate"
type Service struct {
}
func (Service)UserCate() *user_cate.UserCate  {
	return new(user_cate.UserCate)
}
