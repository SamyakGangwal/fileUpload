package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"vicara.co/fileUpload/controller"
	"vicara.co/fileUpload/service"
)

var (
	fileService service.FileService = service.New()
	fileController controller.FileController = controller.New(fileService)
)

func main() {
	server := gin.Default()

	server.GET("/files", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"List of Files" : fileController.FindAll(),
		})
	})

	server.POST("/files", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		fileNames := form.Value["FileName[]"]
		log.Println(files)
		fileController.Save(fileNames, files,ctx)
		ctx.String(http.StatusOK, fmt.Sprintf("%d files were uploaded successfully!", len(files)))
		
	})

	server.Run(":8080")
}