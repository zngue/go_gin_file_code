package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type ApiFile struct {
	ApiPath string
	ApiDetailPath string
	ApiDelAndStatusPath string
	ApiEditPath string
	ApiListPath string

}
func (a *ApiFile) SetInit(create *FileCreate)  {
	a.ApiPath = create.ModelName+".go"
	a.ApiDetailPath = "detail.go"
	a.ApiEditPath = "edit.go"
	a.ApiDelAndStatusPath = "del_or_status.go"
	a.ApiListPath = "list.go"
}
func (a *ApiFile) ApiFileInit(create *FileCreate)  {
	a.SetInit(create)
	a.AliPath(create)
	a.ApiDelAndStatus(create)
	a.ApiDetail(create)
	a.ApiEdit(create)
	a.ApiList(create)
}
func (a *ApiFile ) AliPath(create *FileCreate) error {

	tep := create.TempFile.ApiPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.ApiFilePath+"/"+a.ApiPath //文件路径
	f,ers:=os.Create(filePath)
	if ers!=nil {
		return ers
	}
	str :=string(fByte)
	modelName :=new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str,"{{model}}",create.ModelName)
	str = strings.ReplaceAll(str,"{{GoModelPath}}",create.GoModPath)
	str = strings.ReplaceAll(str,"{{StructName}}",modelName)
	str = strings.ReplaceAll(str,"{{serviceFile}}",create.ModelsPath.ServiceModelPath)
	_,err:=f.Write([]byte(str))
	if err!=nil {
		return err
	}
	return nil
}
func (a *ApiFile ) ApiDelAndStatus(create *FileCreate) error {
	tep := create.TempFile.ApiDelOrStatusPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.ApiFilePath+"/"+a.ApiDelAndStatusPath //文件路径
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
func (a *ApiFile ) ApiDetail(create *FileCreate) error {
	tep := create.TempFile.ApiDetailPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.ApiFilePath+"/"+a.ApiDetailPath //文件路径
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
func (a *ApiFile ) ApiEdit(create *FileCreate) error {
	tep := create.TempFile.ApiEditPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.ApiFilePath+"/"+a.ApiEditPath //文件路径
	f,ers:=os.Create(filePath)
	if ers!=nil {
		return ers
	}
	str :=string(fByte)
	modelName :=new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str,"{{model}}",create.ModelName)
	str = strings.ReplaceAll(str,"{{GoModelPath}}",create.GoModPath)
	str = strings.ReplaceAll(str,"{{StructName}}",modelName)
	str = strings.ReplaceAll(str,"{{modelFile}}",create.ModelsPath.ModelsPath)
	_,err:=f.Write([]byte(str))
	if err!=nil {
		return err
	}
	return nil
}
func (a *ApiFile ) ApiList(create *FileCreate) error {
	tep := create.TempFile.ApiListPath//获取模板信息
	fByte,errs:=ioutil.ReadFile(tep)
	if errs!=nil {
		return errs
	}
	filePath:=create.BaseFile.ApiFilePath+"/"+a.ApiListPath //文件路径
	f,ers:=os.Create(filePath)
	if ers!=nil {
		return ers
	}
	str :=string(fByte)
	str = strings.ReplaceAll(str,"{{model}}",create.ModelName)
	str = strings.ReplaceAll(str,"{{GoModelPath}}",create.GoModPath)
	str = strings.ReplaceAll(str,"{{requestFile}}",create.ModelsPath.RequestModelPath)
	_,err:=f.Write([]byte(str))
	if err!=nil {
		return err
	}
	return nil
}


