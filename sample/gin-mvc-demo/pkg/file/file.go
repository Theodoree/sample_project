package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 返回拓展名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//查看文件是否存在
func CheckExist(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}

//检查文件权限
func CheckPermission(path string) bool {
	_, err := os.Stat(path)

	return os.IsPermission(err)
}

//如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//MkDir
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd() //获取当前路径
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath  //获取文件路径
	perm := CheckPermission(src) //检查权限
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src) //检查MKdir
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644) //打开文件,如果不存在则创建
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
