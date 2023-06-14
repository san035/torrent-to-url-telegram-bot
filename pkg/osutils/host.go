package osutils

import (
	"fmt"
	"syscall"
)

func GetFreeMem() string {

	var r syscall.Rusage
	err := syscall.Getrusage(syscall.RUSAGE_SELF, &r)
	if err != nil {
		return ""
	}
	freeMem := r.Maxrss / 1024 / 1024 // свободная память в Гб
	return fmt.Sprintf("Free memory: %d Gb", freeMem)
}

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
	free := uint64(stat.Bsize) * stat.Bfree / 1024 / 1024 / 1024 // свободное место на диске в Гб
	return fmt.Sprintf("Free HDD: %d Gb", free)
}
