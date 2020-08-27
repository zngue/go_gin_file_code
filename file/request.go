package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type RequestFile struct {
	RequestList string //实体页面
	
}
func (r *RequestFile) SetInit(create *FileCreate) {
	r.RequestList = create.ModelName + ".go"
}
func ( r *RequestFile) RequestFileInit(create *FileCreate)  {
	r.SetInit(create)
	r.CreateRequestList(create)
}
func (r *RequestFile) CreateRequestList(create *FileCreate) error {

	tep := create.TempFile.RequestPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.RequestFilePath+"/"+r.RequestList
	f,ers:=os.Create(filePath)
	if ers!=nil {
		return ers
	}
	str :=string(fByte)
	str = strings.ReplaceAll(str,"{{model}}",create.ModelName)
	_,err:=f.Write([]byte(str))
	if err!=nil {
		return err
	}
	return nil

}