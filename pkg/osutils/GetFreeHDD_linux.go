package osutils

import (
	"fmt"
	"syscall"
)

func GetFreeHDD() string {
	var stat syscall.Statfs_t
	wd, err := syscall.Getwd()
	if err != nil {
		return ""
	}
	err = syscall.Statfs(wd, &stat)
	if err != nil {
		return ""
	}
	free := float32(uint64(stat.Bsize)*stat.Bfree) / (1024 * 1024 * 1024) // свободное место на диске в Гб
	return fmt.Sprintf("Free HDD: %.2f Gb", free)
}
