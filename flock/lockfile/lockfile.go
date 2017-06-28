package lockfile

import "os"

//文件锁
type FileLock struct {
	dir string
	f   *os.File
}

func NewFileLock(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}
