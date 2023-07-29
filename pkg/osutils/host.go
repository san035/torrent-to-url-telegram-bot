package osutils

import (
	"fmt"
	"main.go/internal/web_server"
	"runtime"
	"strconv"
)

const MEGABYTE = 1024 * 1024

func GetMem() string {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return fmt.Sprintf("Free memory: %.2f Mb\nAllocated memory: %.2f Mb\nTotal memory: %.2f Mb",
		float32(m.Frees)/MEGABYTE, float32(m.Alloc)/MEGABYTE, float32(m.Sys)/MEGABYTE)
}

func InfoHost() string {
	return "host: " +
		*web_server.HostAndPort +
		"\n" + GetFreeHDD() +
		"\n" + GetMem() +
		"\n Goroutines: " + strconv.Itoa(runtime.NumGoroutine())
}
