package osutils

import (
	"fmt"
	"syscall"
	"unsafe"
)

func GetFreeHDD() string {
	var freeBytes int64
	disk := "C:"
	_, _, _ = syscall.NewLazyDLL("kernel32.dll").NewProc("GetDiskFreeSpaceExW").Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(disk))),
		uintptr(unsafe.Pointer(&freeBytes)),
		0,
		0,
	)

	freeSpace := float64(freeBytes) / (1024 * 1024 * 1024) // Конвертируем в гигабайты

	return fmt.Sprintf("Free HDD %s %.2f GB", disk, freeSpace)
}
