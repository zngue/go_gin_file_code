package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type ServiceFile struct {
	ServicePathDetail    string //详情页面
	ServicePath          string //实体页面
	ServicePathEdit      string //编辑接口
	ServicePathStatusDel string //修改状态或者删除
	ServicePathList      string //列表数据
}

func (s *ServiceFile) SetInit(create *FileCreate) {
	s.ServicePath = create.ModelName + ".go"
	s.ServicePathDetail = "detail.go"
	s.ServicePathEdit = "edit.go"
	s.ServicePathStatusDel = "del_or_status.go"
	s.ServicePathList = "list.go"
}
func (s *ServiceFile) ServiceFileInit(create *FileCreate) {
	s.SetInit(create)
	s.Service(create)
	s.ServiceDetail(create)
	s.ServiceEdit(create)
	s.ServiceStatusDel(create)
	s.ServiceList(create)
}

func (s *ServiceFile) Service(create *FileCreate) error {
	tep := create.TempFile.ServicePath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ServiceFilePath + "/" + s.ServicePath
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{modelFile}}", create.ModelsPath.ModelsPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceFile) ServiceDetail(create *FileCreate) error {
	tep := create.TempFile.ServiceDetailPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ServiceFilePath + "/" + s.ServicePathDetail
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{modelFile}}", create.ModelsPath.ModelsPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceFile) ServiceEdit(create *FileCreate) error {
	tep := create.TempFile.ServiceEditPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ServiceFilePath + "/" + s.ServicePathEdit
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{modelFile}}", create.ModelsPath.ModelsPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceFile) ServiceStatusDel(create *FileCreate) error {
	tep := create.TempFile.ServiceStatusDelPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ServiceFilePath + "/" + s.ServicePathStatusDel
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{modelFile}}", create.ModelsPath.ModelsPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceFile) ServiceList(create *FileCreate) error {
	tep := create.TempFile.ServiceListPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ServiceFilePath + "/" + s.ServicePathList
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{modelFile}}", create.ModelsPath.ModelsPath)
	str = strings.ReplaceAll(str, "{{requestFile}}", create.ModelsPath.RequestModelPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
