package upload

import (
	"Aoi/global"
	"Aoi/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

const MB = 1024 * 1024

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 获取哈希值
func GetFileName(name string) string {
	ext := getFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//带dot点
func getFileExt(value string) string {
	return path.Ext(value)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, fileName string) bool {
	ext := getFileExt(fileName) //获取到尾缀
	//同意转化为大写比较
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}
	return false
}

// CheckMaxSize 检验文件大小是否合法
func CheckMaxSize(t FileType, f *multipart.FileHeader) bool {
	//all, _ := ioutil.ReadAll(f)
	//size := len(all)
	size := int(f.Size)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*MB {
			return false
		}
	}
	return true
}

// CheckPermission 查看权限是否充足
func CheckPermission(path string) bool {
	_, stat := os.Stat(path)
	return os.IsPermission(stat)
}

// CreateSavaPath 创建文件的保存路径
func CreateSavaPath(path string, perm os.FileMode) error {
	savePath := CheckSavePath(path)
	if savePath {
		err := os.MkdirAll(path, perm)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveFile 将文件保存至指定地址
func SaveFile(file *multipart.FileHeader, path string) error {
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer create.Close()
	_, err = io.Copy(create, open)
	return err
}
