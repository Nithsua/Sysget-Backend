package top

import (
	"os"
	"syscall"
)

//GetDiskUsage ...
func (top *SystemTopStruct) GetDiskUsage() error {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	syscall.Statfs(wd, &stat)
	top.totalDiskSpace = stat.Blocks * uint64(stat.Bsize)
	top.diskSpaceUsed = (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)
	top.diskSpaceFree = stat.Bfree * uint64(stat.Bsize)
	return nil
}
