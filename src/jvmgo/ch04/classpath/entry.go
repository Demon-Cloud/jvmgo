package classpath

import (
	"os"
	"strings"
)

// 操作系统指定的表分隔符
const pathListSeparator = string(os.PathListSeparator)

// 类路径接口定义
type Entry interface {
	// 查找加载 class 文件, className是相对路径
	readClass(className string) ([]byte, Entry, error)

	// 返回变量的字符串
	String() string
}

// 根据 path 创建不同的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,  "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || 
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}