package service

import (
	"fmt"
	"path"
    "path/filepath"
    "runtime"
	"os"
	"vicara.co/fileUpload/entity"
	"github.com/gin-gonic/gin"
)

// Working directory
var mydir = rootDir()

// FileService interface defiing the list of functions needed
type FileService interface {
	Save(file entity.UploadFiles, ctx *gin.Context) entity.UploadFiles
	FindAll() []string
}

// fileService struct
type fileService struct {
	files []entity.UploadFiles
}
// New initiliazer 
func New() FileService {
	return &fileService{}
}


func (service *fileService) Save(file entity.UploadFiles, ctx *gin.Context) entity.UploadFiles {
	// service.files = append(service.files, file)
	
	
	fmt.Println(mydir)
	
	fileList := file.FileData

	for i, f := range fileList {
		
		ctx.SaveUploadedFile(f, mydir + "\\fileStore\\" + file.FileName[i])
	}

	
	return file
}

func (service *fileService) FindAll() []string {

	var files []string

    
    err := filepath.Walk(mydir + "\\fileStore", func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }

	files = files[1:]

	return files
}

func rootDir() string {
    _, b, _, _ := runtime.Caller(0)
    d := path.Join(path.Dir(b))
    return filepath.Dir(d)
}