package file

type FileCreate struct {
	ModelName string
	TempBasePath string
	GoModPath  string
	BaseFile
	TempFile
	ModelFile
	ServiceFile
	RequestFile
	ApiFile
	RouterFile
}
func (f *FileCreate) FileCreateInit()  {//初始化数据
	f.BaseFileInit(f.ModelName)
	f.TempFile.TempFileInit(f)
	f.ModelFile.CreateModel(f)
	f.ServiceFile.ServiceFileInit(f)
	f.RequestFile.RequestFileInit(f)
	f.ApiFile.ApiFileInit(f)
	f.RouterFile.RouterFileInit(f)
}

















