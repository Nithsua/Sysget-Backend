package top

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

//CPUCoreTop is used to hold information per clock
type CPUCoreTop struct {
	CoreUtil float64
}

//getCPUProcData reads /proc/stat and returns total and idle time
func getCPUProcData() ([]byte, error) {
	data, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		log.Fatal("Error opening the stat", err.Error())
		return data, err
	}
	return data, nil
}

//CalculateCPUUtil calculates CPU utilization with delta of 1 second
func (top *SystemTopStruct) CalculateCPUUtil() error {
	var totalTime, idleTime []int64
	for i := 0; i < 2; i++ {
		procData, err := getCPUProcData()
		if err != nil {
			log.Fatal("")
			return err
		}
		stringProcData := strings.Split(string(procData), "\n")
		top.cpuCoreCount = getCoreCount(stringProcData)
		// fmt.Println(stringProcData)
		for j := 0; j < top.cpuCoreCount+1; j++ {
			tempTotalTime, tempIdleTime, err := calculateIdleAndTotalTime(stringProcData[j])
			if err != nil {
				log.Fatal("")
				return err
			}
			if i == 0 {
				totalTime = append(totalTime, tempTotalTime)
				idleTime = append(idleTime, tempIdleTime)
			} else {
				if j == 0 {
					deltaTotalTime := tempTotalTime - totalTime[j]
					deltaIdleTime := tempIdleTime - idleTime[j]
					util, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (1-(float64(deltaIdleTime)/float64(deltaTotalTime)))*100), 64)
					if err != nil {
						return err
					}
					top.cpuPackageUtil = CPUCoreTop{
						CoreUtil: util,
					}
				} else {
					deltaTotalTime := tempTotalTime - totalTime[j]
					deltaIdleTime := tempIdleTime - idleTime[j]
					util, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (1-(float64(deltaIdleTime)/float64(deltaTotalTime)))*100), 64)
					if err != nil {
						return err
					}
					top.perCore = append(top.perCore, CPUCoreTop{
						CoreUtil: util,
					})
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	// fmt.Println(totalTime, idleTime)
	return nil
}

func getCoreCount(stringProcData []string) int {
	coreCount := -1
	for _, v := range stringProcData {
		if strings.Contains(v, "cpu") {
			coreCount++
		}
	}
	return coreCount
}

func calculateIdleAndTotalTime(cpuData string) (int64, int64, error) {
	tempData := strings.Fields(cpuData)
	// fmt.Println(tempData)
	totalTime := int64(0)
	for i := 1; i < len(tempData); i++ {
		temp, err := strconv.ParseInt(tempData[i], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
			return 0, 0, err
		}
		totalTime += temp
	}
	// fmt.Println(tempData)
	idleTime, err := strconv.ParseInt(tempData[4], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return totalTime, idleTime, nil
}
