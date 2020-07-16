package top

import (
	"fmt"
	"log"
)

//Top is used as structure for system usage info
type Top struct {
	cpuTop    CPUTop
	memoryTop MemoryTop
	diskTop   DiskTop
}

//Reset is used to reset the value of the field in SystemTopStruct
func (top *Top) Reset() {
	top.cpuTop = CPUTop{}
	top.memoryTop = MemoryTop{}
	top.diskTop = DiskTop{}
}

//RetriveInfo retrives all the system top info
func (top *Top) RetriveInfo() error {
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
	// err = top.GetDiskInfo()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return err
	// }
	return nil
}

//GetTop returns the entire structure
func (top *Top) GetTop() interface{} {
	return struct {
		CPUTop    CPUTop
		MemoryTop MemoryTop
		DiskTop   DiskTop
	}{CPUTop: top.cpuTop, MemoryTop: top.memoryTop, DiskTop: top.diskTop}
}

//PrintData prints the content of the structure
func (top Top) PrintData() {
	fmt.Printf("CPU Package Utilization: %.2f %% \n", top.cpuTop.CPUPackageUtil.CoreUtil)
	fmt.Println("Per Core Utilization:", top.cpuTop.PerCoreUtil)
	fmt.Println("Memory Info:")
	fmt.Println("Total Memory:", top.memoryTop.MemTotal)
	fmt.Println("Memory Available:", top.memoryTop.MemAvail)
	fmt.Println("Memory Free:", top.memoryTop.MemFree)
	fmt.Println("Disk Info: ")
	fmt.Println(top.diskTop.diskList)
}
