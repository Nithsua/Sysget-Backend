package top

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//CPUCoreTop is used to hold information per clock
type CPUCoreTop struct {
	CoreUtil  float64
	TotalTime uint64
	IdleTime  uint64
}

//CalculateCPUUtil calculates CPU utilization with delta of 1 second
func (top *SystemTopStruct) CalculateCPUUtil() error {
	deltaTT, deltaIT := 0.0, 0.0
	prevTT, prevIT := 0.0, 0.0

	for i := 0; i < 2; i++ {
		procData, err := getProcData()
		if err != nil {
			return err
		}
		totalTime, idleTime, err := calculateIdleAndTotalTime(procData)
		// fmt.Println(totalTime, idleTime)
		if err != nil {
			os.Exit(1)
		} else {
			deltaTT = float64(totalTime) - prevTT
			deltaIT = float64(idleTime) - prevIT
		}
		prevIT = float64(idleTime)
		prevTT = float64(totalTime)
		time.Sleep(1 * time.Second)
	}
	temp := fmt.Sprintf("%.2f", (1-(deltaIT/deltaTT))*100)
	cpuUtil, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return err
	}
	top.cpuPackageUtil.CoreUtil = cpuUtil
	// fmt.Println(*top)
	return nil
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
