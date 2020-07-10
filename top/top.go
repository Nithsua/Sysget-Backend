package top

import (
	"fmt"
	"log"
)

//SystemTopStruct is used as structure for system usage info
type SystemTopStruct struct {
	cpuPackageUtil CPUCoreTop
	cpuCoreCount   int
	perCore        []CPUCoreTop
	memory         Memory
	diskList       []DiskInfo
}

//Reset is used to reset the value of the field in SystemTopStruct
func (top *SystemTopStruct) Reset() {
	top.cpuPackageUtil = CPUCoreTop{}
	top.perCore = []CPUCoreTop{}
	top.cpuCoreCount = 0
	top.diskList = []DiskInfo{}
}

//RetriveInfo retrives all the system top info
func (top *SystemTopStruct) RetriveInfo() error {
	err := top.CalculateCPUUtil()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = top.GetMemoryInfo()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = top.GetDiskData()
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
		Memory         Memory
		DiskList       []DiskInfo
	}{CPUPackageUtil: top.cpuPackageUtil, PerCore: top.perCore, CPUCoreCount: top.cpuCoreCount,
		DiskList: top.diskList, Memory: top.memory}
}

//PrintData prints the content of the structure
func (top SystemTopStruct) PrintData() {
	fmt.Printf("CPU Package Utilization: %.2f %% \n", top.cpuPackageUtil.CoreUtil)
	fmt.Println("Per Core Utilization:", top.perCore)
	fmt.Println("Memory Info:")
	fmt.Println("Total Memory:", top.memory.MemTotal)
	fmt.Println("Memory Available:", top.memory.MemAvail)
	fmt.Println("Memory Free:", top.memory.MemFree)
	fmt.Println("Disk Info:")
	fmt.Println("Total Disk Space: ", top.diskList[0].TotalDiskSpace)
	fmt.Println("Disk Space Used: ", top.diskList[0].DiskSpaceUsed)
	fmt.Println("Disk Space Free:", top.diskList[0].DiskSpaceFree)
}
