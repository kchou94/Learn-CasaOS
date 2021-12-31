package config

import "Learn-CasaOS/pkg/utils/file"

//检查目录是否存在
func mkdirDATAALL() {
	dirArray := [7]string{"/DATA/AppData", "/DATA/Documents", "/DATA/Downloads", "/DATA/Gallery", "/DATA/Media/Movies", "/DATA/Media/TV Shows", "/DATA/Media/Music"}
	for _, dir := range dirArray {
		file.IsNotExistMkDir(dir)
	}
}

func UpdateSetup() {
	mkdirDATAALL()
}
