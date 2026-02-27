package pathPg

import (
	"fmt"
	"os"
)

// DirectoryCreate 创建目录，支持递归创建，并设置适当的权限
func DirectoryCreate(path string) error {
	// 检查目录是否已存在
	if _, err := os.Stat(path); err == nil {
		// 目录已存在，检查是否为目录
		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("检查路径是否为目录失败: %v", err)
		}
		if !info.IsDir() {
			return fmt.Errorf("路径 %s 已存在但不是目录", path)
		}
		return nil
	} else if !os.IsNotExist(err) {
		// 发生其他错误
		return fmt.Errorf("检查目录是否存在失败: %v", err)
	}

	// 目录不存在，创建目录
	// 0755 表示目录权限为 rwxr-xr-x
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 确保目录权限正确设置（某些系统可能需要额外设置）
	err = os.Chmod(path, 0755)
	if err != nil {
		return fmt.Errorf("设置目录权限失败: %v", err)
	}

	return nil
}
