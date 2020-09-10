package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type RouterFile struct {
	RouterPath    string
	RouterApiPath string
}

func (r *RouterFile) SetInit(create *FileCreate) {
	r.RouterPath = create.ModelName + ".go"
	r.RouterApiPath = "router.go"
}
func (r *RouterFile) RouterFileInit(create *FileCreate) {
	r.SetInit(create)
	r.RouterFilePath(create)
	r.RouterApi(create)
}
func (r *RouterFile) RouterFilePath(create *FileCreate) error {
	tep := create.TempFile.RouterPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.RouterFilePath + "/" + r.RouterPath //文件路径
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	str = strings.ReplaceAll(str, "{{GoModelPath}}", create.GoModPath)
	str = strings.ReplaceAll(str, "{{StructName}}", modelName)
	str = strings.ReplaceAll(str, "{{apiFile}}", create.ModelsPath.ApiModelPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (r *RouterFile) RouterApi(create *FileCreate) error {
	filePath := create.BaseFile.RouterFilePath + "/" + r.RouterApiPath //文件路径
	b, err := r.PathExists(filePath)
	if err != nil {
		return err
	}
	if !b {
		r.CreateApi(create)
	}
	byteStr, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str := string(byteStr)
	funName := modelName + "Router(g)"
	if !strings.Contains(str, funName) {
		src := strings.Split(str, "//")
		stres := src[0] + "// \r\n\t" + funName + src[1]
		f, ers := os.Create(filePath)
		if ers != nil {
			return ers
		}
		f.Write([]byte(stres))
	}
	return nil

}
func (r *RouterFile) CreateApi(create *FileCreate) error {
	tep := create.TempFile.RouterApiPath //获取模板信息
	fByte, errs := ioutil.ReadFile(tep)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.RouterFilePath + "/" + r.RouterApiPath //文件路径
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := strings.ReplaceAll(string(fByte), "{{GoModelPath}}", create.GoModPath)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
func (RouterFile) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
