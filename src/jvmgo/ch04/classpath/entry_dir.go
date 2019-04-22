package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录类型的类路径
type DirEntry struct {
	// 目录的绝对路径
	absDir string
}

// new 开头的都为构造函数
func newDirEntry(path string) *DirEntry {
	// 将路径转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 终止程序
		panic(err)
	}
	return &DirEntry{absDir}
}

// 查找读取class
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 将路径拼接完整
	fileName := filepath.Join(self.absDir, className)
	
	// 读取文件字节流
	data, err := ioutil.ReadFile(fileName)

	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}