package logging

import (
	"fmt"
	"go_learning/gin_example/pkg/file"
	"go_learning/gin_example/pkg/setting"
	"os"
	"time"
)

func getLogFilepath() string {
	return fmt.Sprintf("%s%s",
		setting.AppSetting.RuntimeRootPath,
		setting.AppSetting.LogSavePath,
	)
}

func getLogFilename() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func openLogFile(filename, filepath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v\n", err)
	}

	src := dir + "/" + filepath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s\n", src)
	}

	err = file.IsNotExistMkdir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkdir src: %s, err: %v\n", src, err)
	}

	f, err := file.MustOpen(src+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to openfile: %v", err)
	}
	return f, nil

}
