package entity

import "mime/multipart"

// UploadFiles structure to define the input form data
type UploadFiles struct {
	FileName[] string `form:"fileName[]"`
	FileData[] *multipart.FileHeader `form:"upload[]"`
}