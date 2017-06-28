package lockfile


//加锁
func (l *FileLock) Lock() error {

	file, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = file
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}
	defer syscall.FreeLibrary(h)

	addr, err := syscall.GetProcAddress(h, "LockFile")
	if err != nil {
		return err
	}
	for {
		r0, _, _ := syscall.Syscall6(addr, 5, file.Fd(), 0, 0, 0, 1, 0)
		if 0 != int(r0) {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return nil
}
