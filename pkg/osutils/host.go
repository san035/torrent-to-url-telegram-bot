package osutils

import (
	"fmt"
	"main.go/internal/web_server"
	"runtime"
	"strconv"
	"syscall"
)

const MEGABYTE = 1024 * 1024

func GetMem() string {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return fmt.Sprintf("Free memory: %.2f Mb\nAllocated memory: %.2f Mb\nTotal memory: %.2f Mb",
		float32(m.Frees)/MEGABYTE, float32(m.Alloc)/MEGABYTE, float32(m.Sys)/MEGABYTE)
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
	free := float32(uint64(stat.Bsize)*stat.Bfree) / 1024 / MEGABYTE // свободное место на диске в Гб
	return fmt.Sprintf("Free HDD: %.2f Gb", free)
}

func InfoHost() string {
	return "host: " +
		*web_server.HostAndPort +
		"\n" + GetFreeHDD() +
		"\n" + GetMem() +
		"\n Goroutines: " + strconv.Itoa(runtime.NumGoroutine())
}
