package file

import (
	"os"
	"sync"
)

type CreatePath struct {
	ModelName string
	BasePath string
	ApiPath string
	RequestPath string
	ServicePath string
	ModelPath string
	RouterPath string
	MigrationPath string
}
type FilePath struct {
	ApiFilePath string
	RequestFilePath string
	ServiceFilePath string
	ModelFilePath string
	RouterFilePath string
	MigrationFilePath string
}


type BaseFile struct {
	CreatePath
	FilePath
	ModelsPath
}

func (b *BaseFile) BaseFileInit(ModelName ,ProjectName string)  {
	b.ModelName = ModelName
	b.BasePath="app"
	if b.ApiPath==""{
		b.ApiPath="api"
	}
	if b.RequestPath=="" {
		b.RequestPath="request"
	}
	if b.ServicePath=="" {
		b.ServicePath="service"
	}
	if b.ModelPath=="" {
		b.ModelPath="model"
	}
	if b.RouterPath=="" {
		b.RouterPath="router"
	}
	if b.MigrationPath=="" {
		b.MigrationPath="migration"
	}
	if ProjectName!="" {
		b.FilePath.ApiFilePath= ProjectName+"/"+b.CreatePath.BasePath+"/"+ b.CreatePath.ApiPath+"/"+ b.ModelName
		b.FilePath.RequestFilePath= ProjectName+"/"+b.CreatePath.BasePath+"/"+ b.CreatePath.RequestPath+"/"+ b.ModelName
		b.FilePath.ServiceFilePath= ProjectName+"/"+b.CreatePath.BasePath+"/"+ b.CreatePath.ServicePath+"/"+ b.ModelName
		b.FilePath.ModelFilePath= ProjectName+"/"+b.CreatePath.BasePath+"/"+ b.CreatePath.ModelPath+"/"+ b.ModelName
		b.FilePath.RouterFilePath= ProjectName+"/"+b.RouterPath
		b.FilePath.MigrationFilePath= ProjectName+"/"+b.MigrationPath
	}else{
		b.FilePath.ApiFilePath= b.CreatePath.BasePath+"/"+ b.CreatePath.ApiPath+"/"+ b.ModelName
		b.FilePath.RequestFilePath= b.CreatePath.BasePath+"/"+ b.CreatePath.RequestPath+"/"+ b.ModelName
		b.FilePath.ServiceFilePath= b.CreatePath.BasePath+"/"+ b.CreatePath.ServicePath+"/"+ b.ModelName
		b.FilePath.ModelFilePath= b.CreatePath.BasePath+"/"+ b.CreatePath.ModelPath+"/"+ b.ModelName
		b.FilePath.RouterFilePath= b.RouterPath
		b.FilePath.MigrationFilePath= b.MigrationPath
	}
	b.CreateInit()
	b.ModelsPathInit()
}
type ModelsPath struct {
	ApiModelPath       string
	RequestModelPath   string
	ServiceModelPath   string
	ModelsPath         string
	RouterModelPath    string
	MigrationModelPath string
}
func (b *BaseFile) ModelsPathInit()  {
	b.ModelsPath.ApiModelPath= b.CreatePath.BasePath+"/"+ b.CreatePath.ApiPath+"/"+ b.ModelName
	b.ModelsPath.RequestModelPath= b.CreatePath.BasePath+"/"+ b.CreatePath.RequestPath+"/"+ b.ModelName
	b.ModelsPath.ServiceModelPath= b.CreatePath.BasePath+"/"+ b.CreatePath.ServicePath+"/"+ b.ModelName
	b.ModelsPath.ModelsPath= b.CreatePath.BasePath+"/"+ b.CreatePath.ModelPath+"/"+ b.ModelName
	b.ModelsPath.RouterModelPath= b.RouterPath
	b.ModelsPath.MigrationModelPath=b.MigrationPath
}
//创建文件夹

func (b *BaseFile) CreateInit()  {
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.ApiFilePath)
	}()
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.RequestFilePath)
	}()
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.ServiceFilePath)
	}()
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.ModelFilePath)
	}()
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.RouterFilePath)
	}()
	go func() {
		defer wg.Done()
		b.CreateMutiDir(b.MigrationFilePath)
	}()
	wg.Wait()
}
//调用os.MkdirAll递归创建文件夹
func (b *BaseFile) CreateMutiDir(filePath string) error {
	if !b.isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func (b *BaseFile) isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
