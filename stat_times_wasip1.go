//go:build wasip1

package copy

import (
	"os"
	"syscall"
	"time"
)

func getTimeSpec(info os.FileInfo) timespec {
	stat := info.Sys().(*syscall.Stat_t)
	times := timespec{
		Mtime: info.ModTime(),
		Atime: time.Unix(int64(stat.Atime), 0),
		Ctime: time.Unix(int64(stat.Ctime), 0),
	}
	return times
}
