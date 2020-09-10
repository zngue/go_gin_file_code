package file

type TempFile struct {
	BasePath string
	ApiTemp
	ModelTemp
	ServiceTemp
	RequestTemp
	RouterTemp
}

func (t *TempFile) TempFileInit(create *FileCreate) {
	t.BasePath = create.TempBasePath
	t.ModelTempInit(create.TempBasePath)
	t.ServiceTempInit(create)
	t.RequestTempInit(create)
	t.ApiTempInit(create)
	t.RouterTempInit(create)
}

type ModelTemp struct {
	ModelTempPath string
}

func (m *ModelTemp) ModelTempInit(basePath string) {
	m.ModelTempPath = basePath + "/model/model.tpl"
}

type ServiceTemp struct {
	ServicePath          string
	ServiceDetailPath    string
	ServiceEditPath      string
	ServiceStatusDelPath string
	ServiceListPath      string
}

func (s *ServiceTemp) ServiceTempInit(create *FileCreate) {
	s.ServicePath = create.TempBasePath + "/service/service.tpl"
	s.ServiceDetailPath = create.TempBasePath + "/service/detail.tpl"
	s.ServiceEditPath = create.TempBasePath + "/service/edit.tpl"
	s.ServiceListPath = create.TempBasePath + "/service/list.tpl"
	s.ServiceStatusDelPath = create.TempBasePath + "/service/del_or_status.tpl"
}

type RequestTemp struct {
	RequestPath string
}

func (s *RequestTemp) RequestTempInit(create *FileCreate) {
	s.RequestPath = create.TempBasePath + "/request/request.tpl"
}

type ApiTemp struct {
	ApiPath            string
	ApiDelOrStatusPath string
	ApiDetailPath      string
	ApiEditPath        string
	ApiListPath        string
}

func (a *ApiTemp) ApiTempInit(create *FileCreate) {
	a.ApiPath = create.TempBasePath + "/api/api.tpl"
	a.ApiDelOrStatusPath = create.TempBasePath + "/api/del_or_status.tpl"
	a.ApiDetailPath = create.TempBasePath + "/api/detail.tpl"
	a.ApiEditPath = create.TempBasePath + "/api/edit.tpl"
	a.ApiListPath = create.TempBasePath + "/api/list.tpl"
}

type RouterTemp struct {
	RouterPath    string
	RouterApiPath string
}

func (s *RouterTemp) RouterTempInit(create *FileCreate) {
	s.RouterPath = create.TempBasePath + "/router/router.tpl"
	s.RouterApiPath = create.TempBasePath + "/router/api.tpl"
}
