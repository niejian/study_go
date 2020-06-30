package util

import "os"

// 检查文件、目录是否存在
func PathExists(path string) (bool, error)  {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
