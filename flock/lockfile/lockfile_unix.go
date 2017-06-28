package lockfile

import (
	"os"
	"syscall"
	"fmt"
)

// 加锁
func (l *FileLock) Lock() error {
	f, err := os.OpenFile(l.dir,os.O_WRONLY | os.O_CREATE,0664)
	if err != nil {
		return err
	}
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		f.Close() 	// 如果加锁失败，说明有人使用，那么关闭现有的文件
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	l.f = f // 加锁成功，关联文件
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
