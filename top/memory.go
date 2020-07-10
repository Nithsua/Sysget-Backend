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
	Writeback    int64
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
		case strings.Compare(tempList[0], "MemTotal:") == 0:
			tempMemory.MemTotal = temp
		case strings.Compare(tempList[0], "MemFree:") == 0:
			tempMemory.MemFree = temp
		case strings.Compare(tempList[0], "MemAvailable:") == 0:
			tempMemory.MemAvail = temp
		case strings.Compare(tempList[0], "Buffers:") == 0:
			tempMemory.Buffers = temp
		case strings.Compare(tempList[0], "Cached:") == 0:
			tempMemory.Cached = temp
		case strings.Compare(tempList[0], "SwapCached:") == 0:
			tempMemory.SwapCached = temp
		case strings.Compare(tempList[0], "Active:") == 0:
			tempMemory.Active = temp
		case strings.Compare(tempList[0], "Inactive:") == 0:
			tempMemory.InActive = temp
		case strings.Compare(tempList[0], "Active(anon):") == 0:
			tempMemory.ActiveAnon = temp
		case strings.Compare(tempList[0], "Inactive(anon):") == 0:
			tempMemory.InActiveAnon = temp
		case strings.Compare(tempList[0], "Active(file):") == 0:
			tempMemory.ActiveFile = temp
		case strings.Compare(tempList[0], "Inactive(file):") == 0:
			tempMemory.InActiveFile = temp
		case strings.Compare(tempList[0], "Unevictable:") == 0:
			tempMemory.Unevictable = temp
		case strings.Compare(tempList[0], "Mlocked:") == 0:
			tempMemory.Mlocked = temp
		case strings.Compare(tempList[0], "SwapTotal:") == 0:
			tempMemory.InActiveFile = temp
		case strings.Compare(tempList[0], "SwapFree:") == 0:
			tempMemory.SwapFree = temp
		case strings.Compare(tempList[0], "Dirty:") == 0:
			tempMemory.Dirty = temp
		case strings.Compare(tempList[0], "Writeback:") == 0:
			tempMemory.Writeback = temp
		case strings.Compare(tempList[0], "AnonPages:") == 0:
			tempMemory.AnonPages = temp
		case strings.Compare(tempList[0], "Mapped:") == 0:
			tempMemory.Mapped = temp
		case strings.Compare(tempList[0], "Shmem:") == 0:
			tempMemory.SharedMem = temp
		}
	}
	top.memory = tempMemory
	return nil
}
