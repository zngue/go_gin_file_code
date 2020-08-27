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
}
func (c *BaseFile) BaseFileInit(ModelName string)  {
	c.ModelName = ModelName
	c.BasePath="app"
	if c.ApiPath==""{
		c.ApiPath="api"
	}
	if c.RequestPath=="" {
		c.RequestPath="request"
	}
	if c.ServicePath=="" {
		c.ServicePath="service"
	}
	if c.ModelPath=="" {
		c.ModelPath="model"
	}
	if c.RouterPath=="" {
		c.RouterPath="router"
	}
	if c.MigrationPath=="" {
		c.MigrationPath="migration"
	}
	c.FilePath.ApiFilePath= c.CreatePath.BasePath+"/"+ c.CreatePath.ApiPath+"/"+ c.ModelName
	c.FilePath.RequestFilePath= c.CreatePath.BasePath+"/"+ c.CreatePath.RequestPath+"/"+ c.ModelName
	c.FilePath.ServiceFilePath= c.CreatePath.BasePath+"/"+ c.CreatePath.ServicePath+"/"+ c.ModelName
	c.FilePath.ModelFilePath= c.CreatePath.BasePath+"/"+ c.CreatePath.ModelPath+"/"+ c.ModelName
	c.FilePath.RouterFilePath=c.RouterPath
	c.FilePath.MigrationFilePath=c.MigrationPath
	c.CreateInit()
}
//创建文件夹

func (c *BaseFile) CreateInit()  {
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.ApiFilePath)
	}()
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.RequestFilePath)
	}()
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.ServiceFilePath)
	}()
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.ModelFilePath)
	}()
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.RouterFilePath)
	}()
	go func() {
		defer wg.Done()
		c.CreateMutiDir(c.MigrationFilePath)
	}()
	wg.Wait()
}
//调用os.MkdirAll递归创建文件夹
func (c *BaseFile) CreateMutiDir(filePath string) error {
	if !c.isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func (c *BaseFile) isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
