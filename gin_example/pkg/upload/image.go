package upload

import (
	"fmt"
	"go_learning/gin_example/pkg/file"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/pkg/setting"
	"go_learning/gin_example/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	filename := util.EncodeMD5(
		strings.TrimSuffix(name, ext),
	)
	return filename + ext
}

func CheckImageExt(name string) bool {
	ext := file.GetExt(name)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

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
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v\n", err)
	}

	err = file.IsNotExistMkdir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkdir err: %v\n", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Perrsion denied src: %s\n", src)
	}

	return nil
}
