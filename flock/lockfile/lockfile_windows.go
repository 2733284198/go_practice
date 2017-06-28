package lockfile

import (
	"os"
	"syscall"
	"errors"
	"fmt"
)

//加锁
func (l *FileLock) Lock() error {

	file, err := os.OpenFile(l.dir,os.O_WRONLY | os.O_CREATE,0664)
	if err != nil {
		return err
	}

	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		file.Close()
		return err
	}
	defer syscall.FreeLibrary(h)
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365202(v=vs.85).aspx
	addr, err := syscall.GetProcAddress(h, "LockFile")
	if err != nil {
		return err
	}

	_, _, errno := syscall.Syscall6(addr, 5, file.Fd(), 0x00000001, 0, 0, 1, 0)
	if errno != 0 {
			return errors.New(fmt.Sprintf("errno:%v",errno))
	}
	l.f = file
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}
	defer syscall.FreeLibrary(h)
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365715(v=vs.85).aspx
	addr, err := syscall.GetProcAddress(h, "UnlockFile")
	if err != nil {
		return err
	}
	_, _, errno := syscall.Syscall6(addr, 5, l.f.Fd(), 0x00000001, 0, 0, 1, 0)
	if errno != 0 {
		return errors.New(fmt.Sprintf("errno:%v",errno))
	}
	return nil
}
