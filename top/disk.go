package top

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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

func getPartitionData() ([]byte, error) {
	tempData, err := ioutil.ReadFile("/proc/partitions")
	var unFormattedDiskList []string
	if err != nil {
		log.Fatal(err.Error())
		return tempData, err
	}
	tempList := strings.Split(string(tempData), "\n")
	for _, v := range tempList {
		decision, err := regexp.MatchString(`sd[a-z]$`, v)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			if decision {
				unFormattedDiskList = append(unFormattedDiskList, v)
			}
		}
	}

	return _, nil
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
