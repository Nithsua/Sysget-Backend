package top

import (
	"fmt"
	"strconv"
	"strings"
)

//CPUCoreTop is used to hold information per clock
type CPUCoreTop struct {
	CoreUtil  float64
	TotalTime uint64
	IdleTime  uint64
}

func calculateIdleAndTotalTime(procData []byte) (int64, int64, error) {
	sysUtil := strings.Split(string(procData), "\n")

	cpuData := make([]string, 1, 129)

	for i := 0; i < len(sysUtil); i++ {
		if strings.Contains(sysUtil[i], "cpu") {
			cpuData = append(cpuData, sysUtil[i])
		}
	}

	tempData := strings.Split(cpuData[1], " ")
	totalTime := int64(0)
	for i := 2; i < len(tempData); i++ {
		temp, err := strconv.ParseInt(tempData[i], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return 0, 0, err
		}
		totalTime += temp
	}
	idleTime, err := strconv.ParseInt(tempData[5], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return totalTime, idleTime, nil

}
