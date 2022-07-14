package service

import (
	"Aoi/global"
	"Aoi/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType,
	_ multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	//file为文件流。fileHeader记录文件信息
	fileName := upload.GetFileName(fileHeader.Filename)
	savaPath := upload.GetSavePath()
	dst := savaPath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	//不存在返回true
	if upload.CheckSavePath(savaPath) {
		err := upload.CreateSavaPath(savaPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if !upload.CheckMaxSize(fileType, fileHeader) {
		return nil, errors.New("file is to large")
	}
	if upload.CheckPermission(savaPath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil

}
