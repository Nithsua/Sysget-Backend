package top

import (
	"os"
	"syscall"
)

//DiskInfo ...
type DiskInfo struct {
	DiskName       string
	DiskID         string
	TotalDiskSpace uint64
	DiskSpaceUsed  uint64
	DiskSpaceFree  uint64
}

//GetDiskData ...
func (top *SystemTopStruct) GetDiskData() error {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	syscall.Statfs(wd, &stat)
	top.diskList = append(top.diskList,
		DiskInfo{
			TotalDiskSpace: stat.Blocks * uint64(stat.Bsize),
			DiskSpaceUsed:  (stat.Blocks - stat.Bfree) * uint64(stat.Bsize),
			DiskSpaceFree:  stat.Bfree * uint64(stat.Bsize)})
	return nil
}
