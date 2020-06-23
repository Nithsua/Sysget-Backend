package top

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//SystemTopStruct is used as structure for system usage info
type SystemTopStruct struct {
	CPUUtil float64
}

//GetProcData reads /proc/stat and returns total and idle time
func GetProcData() (int64, int64, error) {
	data, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		log.Fatal("Error opening the stat", err.Error())
	}
	sysUtil := strings.Split(string(data), "\n")

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

//GetCPUUtil calculates CPU utilization with delta of 1 second
func GetCPUUtil() (float64, error) {
	deltaTT, deltaIT := 0.0, 0.0
	prevTT, prevIT := 0.0, 0.0

	for i := 0; i < 2; i++ {
		totalTime, idleTime, err := GetProcData()
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
	return strconv.ParseFloat(temp, 64)
}
