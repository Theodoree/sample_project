package api

import (
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/app"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	AppG := app.Gin{c}

	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image") //获取image文件
	if err != nil {
		logging.Warn(err)
		AppG.Response(http.StatusOK, errors.ERROR, data)
		return
	}

	if image == nil {
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, data)
		return
	}
	imageName := upload.GetImageName(image.Filename) //进行一次md5处理,返回文件名
	fullPath := upload.GetImageFullPath()            // 返回本地储存位置
	savePath := upload.GetImagePath()                //访问地址

	paht := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) { //检查后缀和大小
		AppG.Response(http.StatusOK, errors.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, data)
		return
	}
	err = upload.CheckImage(fullPath) //检测文件夹和权限
	if err != nil {
		logging.Warn(err)
		AppG.Response(http.StatusOK, errors.ERROR_UPLOAD_CHECK_IMAGE_FAIL, data)
		return
	}
	if err := c.SaveUploadedFile(image, paht); err != nil { //储存到本地
		logging.Warn(err)
		AppG.Response(http.StatusOK, errors.ERROR_UPLOAD_SAVE_IMAGE_FAIL, data)
		return
	}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName

	AppG.Response(http.StatusOK, errors.SUCCESS, data)
}
