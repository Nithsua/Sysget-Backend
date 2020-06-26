package top

import (
	"fmt"
	"io/ioutil"
	"log"
)

//SystemTopStruct is used as structure for system usage info
type SystemTopStruct struct {
	cpuPackageUtil CPUCoreTop
	perCore        []CPUCoreTop
	totalDiskSpace uint64
	diskSpaceUsed  uint64
	diskSpaceFree  uint64
}

//getProcData reads /proc/stat and returns total and idle time
func getProcData() ([]byte, error) {
	data, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		log.Fatal("Error opening the stat", err.Error())
		return data, err
	}
	return data, nil
}

//RetriveInfo retrives all the system top info
func (top *SystemTopStruct) RetriveInfo() error {
	err := top.CalculateCPUUtil()
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
