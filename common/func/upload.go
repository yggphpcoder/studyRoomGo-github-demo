package _func

import "os"

/**
*此代码主要用于编辑图片时，删除原有图片
 * 判断文件是否存在  存在返回 true 不存在返回false
*/
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
