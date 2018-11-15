package export

import "github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"

//获取访问路径		127.0.0.1:8080/export/pic.jpg
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// export/
func GetExcelPath() string { //获取储存路径
	return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string { //获取完成的储存路径      runtime/export/
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}