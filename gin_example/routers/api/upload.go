package api

import (
	"github.com/gin-gonic/gin"
	"go_learning/gin_example/pkg/e"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/pkg/upload"
	"net/http"
)

func UploadImage(ctx *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]any)

	file, image, err := ctx.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()

		src := fullPath + imageName

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.INVALID_PARAMS
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.INVALID_PARAMS
			} else if err = ctx.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = upload.GetImagePath() + imageName
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
