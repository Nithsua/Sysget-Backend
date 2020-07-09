package top

import (
	"fmt"
	"io/ioutil"
	"log"
)

//SystemTopStruct is used as structure for system usage info
type SystemTopStruct struct {
	cpuPackageUtil CPUCoreTop
	cpuCoreCount   int
	perCore        []CPUCoreTop
	diskList       []DiskTop
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

//Reset is used to reset the value of the field in SystemTopStruct
func (top *SystemTopStruct) Reset() {
	top.cpuPackageUtil = CPUCoreTop{}
	top.perCore = []CPUCoreTop{}
	top.cpuCoreCount = 0
	top.diskList = []DiskTop{}
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
		PerCore        []CPUCoreTop
		CPUCoreCount   int
		DiskList       []DiskTop
	}{CPUPackageUtil: top.cpuPackageUtil, PerCore: top.perCore, CPUCoreCount: top.cpuCoreCount, DiskList: top.diskList}
}

//PrintData prints the content of the structure
func (top SystemTopStruct) PrintData() {
	fmt.Printf("CPU Package Utilization: %.2f %% \n", top.cpuPackageUtil.CoreUtil)
	fmt.Println("Per Core Utilization:", top.perCore)
	fmt.Println("Disk Info: ", top.diskList)
}
