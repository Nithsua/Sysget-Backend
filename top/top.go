package top

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

//SystemTopStruct is used as structure for system usage info
type SystemTopStruct struct {
	cpuPackageUtil CPUCoreTop
	perCore        []CPUCoreTop
	totalDiskSpace uint64
	diskSpaceUsed  uint64
	diskSpaceFree  uint64
}

//CPUCoreTop is used to hold information per clock
type CPUCoreTop struct {
	CoreUtil  float64
	TotalTime uint64
	IdleTime  uint64
}

//GetDiskUsage ...
func (top *SystemTopStruct) GetDiskUsage() error {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	syscall.Statfs(wd, &stat)
	top.totalDiskSpace = stat.Blocks * uint64(stat.Bsize)
	top.diskSpaceUsed = (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)
	top.diskSpaceFree = stat.Bfree * uint64(stat.Bsize)
	return nil
}

//getProcData reads /proc/stat and returns total and idle time
func getProcData() (int64, int64, error) {
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
func (top *SystemTopStruct) GetCPUUtil() error {
	deltaTT, deltaIT := 0.0, 0.0
	prevTT, prevIT := 0.0, 0.0

	for i := 0; i < 2; i++ {
		totalTime, idleTime, err := getProcData()
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

//RetriveInfo retrives all the system top info
func (top *SystemTopStruct) RetriveInfo() error {
	err := top.GetCPUUtil()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = top.GetDiskUsage()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

//GetTop returns the entire structure
func (top SystemTopStruct) GetTop() interface{} {
	return struct {
		CPUPackageUtil CPUCoreTop
		TotalDiskSpace uint64
		DiskSpaceUsed  uint64
		DiskSpaceFree  uint64
	}{CPUPackageUtil: top.cpuPackageUtil, TotalDiskSpace: top.totalDiskSpace, DiskSpaceUsed: top.diskSpaceUsed, DiskSpaceFree: top.diskSpaceFree}
}

//PrintData prints the content of the structure
func (top SystemTopStruct) PrintData() {
	fmt.Printf("CPU Utilization: %.2f %% \n", top.cpuPackageUtil.CoreUtil)
	fmt.Printf("Total Disk Space: %.2f MB\n", (float64(top.totalDiskSpace)/1024)/1024)
	fmt.Printf("Disk Space used: %.2f MB\n", (float64(top.diskSpaceUsed)/1024)/1024)
	fmt.Printf("Disk Space free: %.2f MB\n", (float64(top.diskSpaceFree)/1024)/1024)
}
