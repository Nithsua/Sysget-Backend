package top

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//DiskTop ...
type DiskTop struct {
	diskCount      int
	diskList       []DiskInfo
	totalDiskSpace uint64
	diskSpaceFree  uint64
	diskSpaceUsed  uint64
}

//DiskInfo ...
type DiskInfo struct {
	DiskName      string
	DiskID        string
	PartitionList []PartitionTop
}

//PartitionTop ...
type PartitionTop struct {
	PartitionID    string
	PartitionType  string
	DiskMountPoint string
	TotalDiskSpace uint64
	DiskSpaceUsed  uint64
	DiskSpaceFree  uint64
}

//GetDiskInfo ...
func (top *Top) GetDiskInfo() error {
	diskData, err := getDiskData()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	getDiskCount(diskData)

	return nil
}

func getDiskCount(diskData []string) {
	tempMap := make(map[string]int)
	for _, v := range diskData {
		tempList := strings.Fields(v)
		if _, contains := tempMap[tempList[0]]; contains {
			tempMap[tempList[0]]++
		} else {
			tempMap[tempList[0]] = 1
		}

	}
}

func getDiskData() ([]string, error) {
	partitionData := []string{}
	output, err := exec.Command("dh", "-h").CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
		return partitionData, err
	}
	tempList := strings.Split(string(output), "\n")

	for _, v := range tempList {
		decision, err := regexp.MatchString(`sd([a-z]+)([0-9]+)`, v)
		if err != nil {
			log.Fatal(err.Error())
			return partitionData, err
		}
		if decision {
			partitionData = append(partitionData, v)
		}
	}
	return partitionData, nil
}
