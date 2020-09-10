package file

import "sync"

type FileCreate struct {
	ModelName    string
	TempBasePath string
	GoModPath    string
	ProjectName  string
	BaseFile
	TempFile
	ModelFile
	ServiceFile
	RequestFile
	ApiFile
	RouterFile
}
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(f func()) {
	w.Add(1) //Add必须在go协程外
	go func() {
		defer w.Done()
		f()
	}()
}
func (f *FileCreate) FileCreateInit() { //初始化数据
	f.BaseFileInit(f.ModelName, f.ProjectName)
	f.TempFile.TempFileInit(f)
	f.FileCreate()

}
func (f *FileCreate) FileCreate() {
	wg := new(WaitGroupWrapper)
	wg.Wrap(func() {
		f.ModelFile.CreateModel(f)
	})
	wg.Wrap(func() {
		f.ServiceFile.ServiceFileInit(f)
	})
	wg.Wrap(func() {
		f.RequestFile.RequestFileInit(f)
	})
	wg.Wrap(func() {
		f.ApiFile.ApiFileInit(f)
	})
	wg.Wrap(func() {
		f.RouterFile.RouterFileInit(f)
	})
	wg.Wait()
}
