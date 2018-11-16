package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/file"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
)

//返回文件访问路径
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

//
func GetImageName(name string) string {
	ext := path.Ext(name)                     //获取后缀
	fileName := strings.TrimSuffix(name, ext) //组成全名
	fileName = util.EncodeMD5(fileName)

	return fileName + ext //Md5+ext
}

//返回upload/images/ 路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

//返回本地储存路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

//检测后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

//检测大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd() //获取当前路径
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	//检测文件夹是否存在
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	//检查文件权限
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
