package file

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

func GetExt(filename string) string {
	return path.Ext(filename)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkdir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func Mkdir(src string) error {
	if err := os.MkdirAll(src, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func MustOpen(filename string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
