package top

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Memory ...
type Memory struct {
	MemTotal     int64
	MemFree      int64
	MemAvail     int64
	Buffers      int64
	Cached       int64
	Active       int64
	InActive     int64
	ActiveAnon   int64
	InActiveAnon int64
	ActiveFile   int64
	InActiveFile int64
	Unevictable  int64
	Mlocked      int64
	SwapCached   int64
	SwapTotal    int64
	SwapFree     int64
	Dirty        int64
	AnonPages    int64
	Mapped       int64
	SharedMem    int64
}

func getMemoryProcData() ([]byte, error) {
	memData, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		log.Fatal(err.Error())
		return memData, err
	}
	return memData, nil
}

//GetMemoryInfo ...
func (top *SystemTopStruct) GetMemoryInfo() error {
	tempMemory := Memory{}
	temp, err := getMemoryProcData()
	tempString := string(temp)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	memData := strings.Split(tempString, "\n")
	if memData[len(memData)-1] == "" {
		memData = memData[:len(memData)-2]
	}

	for _, v := range memData {
		tempList := strings.Fields(v)
		temp, err := strconv.ParseInt(tempList[1], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		switch {
		case strings.Contains(v, "MemTotal:"):
			tempMemory.MemTotal = temp
		case strings.Contains(v, "MemFree:"):
			tempMemory.MemFree = temp
		case strings.Contains(v, "MemAvailable:"):
			tempMemory.MemAvail = temp
		}
	}
	top.memory = tempMemory
	return nil
}
