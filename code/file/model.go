package file

import (
	"io/ioutil"
	"os"
	"strings"
)

type ModelFile struct {
}

func (m *ModelFile) CreateModel(create *FileCreate) error {
	te := create.ModelTempPath
	fByte, errs := ioutil.ReadFile(te)
	if errs != nil {
		return errs
	}
	filePath := create.BaseFile.ModelFilePath + "/" + create.ModelName + ".go"
	f, ers := os.Create(filePath)
	if ers != nil {
		return ers
	}
	str := string(fByte)
	modelName := new(FileNameChange).Case2Camel(create.ModelName)
	str = strings.ReplaceAll(str, "{{modelName}}", modelName)
	str = strings.ReplaceAll(str, "{{model}}", create.ModelName)
	_, err := f.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
