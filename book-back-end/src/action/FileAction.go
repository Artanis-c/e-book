package action

import (
	"book-back-end/src/common"
	result "book-back-end/src/domain/models/dto"
	"github.com/gin-gonic/gin"
	"path"
)

const FILE_BASE_PATH = "static/image/"

type FileAction struct {
}

func NewFileAction() *FileAction {
	return &FileAction{}
}

func (c *FileAction) UploadFile(context *gin.Context) result.Result {
	file, err := context.FormFile("file")
	code := common.GenUniqueCode()
	ext := path.Ext(file.Filename)
	newFileName := code + ext
	newFilePath := FILE_BASE_PATH + newFileName
	if err != nil {
		return result.Fail(err.Error())
	}
	if err := context.SaveUploadedFile(file, newFilePath); err != nil {
		return result.Fail(err.Error())
	}
	return result.Success(newFileName)
}
