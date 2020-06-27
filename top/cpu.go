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
	CoreUtil float64
	// TotalTime uint64
	// IdleTime  uint64
}

//CalculateCPUUtil calculates CPU utilization with delta of 1 second
func (top *SystemTopStruct) CalculateCPUUtil() error {
	var deltaTT, deltaIT []float64
	var prevTT, prevIT []float64
	var totalTime, idleTime []int64

	for i := 0; i < 2; i++ {
		procData, err := getProcData()
		if err != nil {
			return err
		}
		sysUtil := strings.Split(string(procData), "\n")

		var cpuData []string

		for i := 0; i < len(sysUtil); i++ {
			if strings.Contains(sysUtil[i], "cpu") {
				cpuData = append(cpuData, sysUtil[i])
			}
		}

		for _, v := range cpuData {
			fmt.Println(v)
		}

		for i := 0; i < len(cpuData); i++ {
			prevTT = append(prevTT, 0.0)
			prevIT = append(prevIT, 0.0)
		}

		for j := 0; j < len(cpuData); j++ {
			tempTotalTime, tempIdleTime, err := calculateIdleAndTotalTime(cpuData[j])
			if err != nil {
				os.Exit(1)
			} else {
				if i == 0 {
					totalTime = append(totalTime, tempTotalTime)
					idleTime = append(idleTime, tempIdleTime)
					deltaTT = append(deltaTT, float64(totalTime[j])-prevTT[j])
					deltaIT = append(deltaIT, float64(idleTime[j])-prevIT[j])
				} else {
					totalTime[j] = tempTotalTime
					idleTime[j] = tempIdleTime
					deltaTT[j] = float64(totalTime[j]) - prevTT[j]
					deltaIT[j] = float64(idleTime[j]) - prevIT[j]
				}
			}
			prevIT = append(prevIT, float64(idleTime[j]))
			prevTT = append(prevTT, float64(totalTime[j]))
		}
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < len(deltaTT); i++ {
		temp := fmt.Sprintf("%.2f", (1-(deltaIT[i]/deltaTT[i]))*100)
		cpuUtil, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			return err
		}
		if i == 0 {
			top.cpuPackageUtil.CoreUtil = cpuUtil
		} else {
			temp := CPUCoreTop{CoreUtil: cpuUtil}
			top.perCore = append(top.perCore, temp)
		}
	}
	return nil
}

func calculateIdleAndTotalTime(cpuData string) (int64, int64, error) {
	tempData := strings.Split(cpuData, " ")
	// fmt.Println(cpuData)
	totalTime := int64(0)
	for i := 2; i < len(tempData); i++ {
		temp, err := strconv.ParseInt(tempData[i], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return 0, 0, err
		}
		totalTime += temp
	}
	// fmt.Println(tempData)
	idleTime, err := strconv.ParseInt(tempData[5], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return totalTime, idleTime, nil

}
