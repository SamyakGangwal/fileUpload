package controller

import (
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"vicara.co/fileUpload/entity"
	"vicara.co/fileUpload/service"
)

// FileController controller for the project
type FileController interface {
	FindAll() []string
	Save(fileNames[] string, fileData[] *multipart.FileHeader,ctx *gin.Context) entity.UploadFiles
}

type controller struct {
	service service.FileService
}
// New constructor
func New(service service.FileService) FileController {
	return &controller {
		service: service,
	}
}
// FindAll for finding all of the present saved files
func (c *controller) FindAll() []string {
	return c.service.FindAll()
}
func (c *controller) Save(fileNames[] string, fileData[] *multipart.FileHeader,ctx *gin.Context) entity.UploadFiles {
	var file entity.UploadFiles
	file.FileName = fileNames
	file.FileData = fileData
	c.service.Save(file,ctx)
	return file
}