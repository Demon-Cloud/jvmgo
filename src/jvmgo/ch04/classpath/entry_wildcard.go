// 实际上也是 CompositeEntry
package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去掉路径末尾的 *
	baseDir := path[:len(path) - 1] 
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 表示该次调用的path参数指定的目录应被跳过,通配符路径不能递归匹配子目录下的JAR文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 遍历 baseDir 目录下的文件树,对每一个目录和文件调用walkFn方法,包括baseDir自己
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}