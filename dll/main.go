package main

import (
	"syscall"
	"unsafe"
	"log"
)

// Golang下通过syscall调用win32的api
// http://wendal.net/2013/0406.html
// Calling a Windows DLL
// https://github.com/golang/go/wiki/WindowsDLLs
// Windows平台Go调用DLL的坑
// http://www.cnblogs.com/concurrency/p/4170657.html

func main() {
	//首先,准备输入参数, GetDiskFreeSpaceEx需要4个参数, 可查MSDN
	dir := "C:"
	lpFreeBytesAvailable := int64(0) //注意类型需要跟API的类型相符
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)

	//获取方法的引用
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		log.Println(err)
		return
	}
	defer syscall.FreeLibrary(kernel32)

	/*
		BOOL GetDiskFreeSpaceEx(
	 		LPCWSTR lpDirectoryName,
			PULARGE_INTEGER lpFreeBytesAvailableToCaller,
			PULARGE_INTEGER lpTotalNumberOfBytes,
			PULARGE_INTEGER lpTotalNumberOfFreeBytes
		);
		*/
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	// 执行之. 因为有4个参数,故取Syscall6才能放得下. 最后2个参数,自然就是0了
	r, _, errno := syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dir))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	if errno != 0 {
		log.Println(errno)
		return
	}

	// 注意, errno并非error接口的, 不可能是nil
	// 而且,根据MSDN的说明,返回值为0就fail, 不为0就是成功
	if r != 0 {
		log.Printf("Free %dmb", lpTotalNumberOfFreeBytes/1024/1024)
	}
}