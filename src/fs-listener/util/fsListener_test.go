package util

import "testing"

func TestGetFsChange(t *testing.T) {
	t.Run("文件变化监控", func(t *testing.T) {
		GetFsChange("/Users/a/logs/demo-muti-registry-producer/")
	})
}
