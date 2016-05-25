package conutil

import (
    "github.com/shirou/gopsutil/cpu"    
	"encoding/json"
	"strconv"
	"strings"
    "github.com/washraf/runcbm/common"   
    "github.com/washraf/runcbm/common/config"
)

//GetClientUtilization ...
func GetContainerUtilization(ContainerId string) (ContainerMetrics , error) {
    res:= ContainerMetrics{
        
    }
    _,ct,err:=getCPUUtlizationPercent(ContainerId)
    if(err!=nil){
        return res,err
    }
    res.CPUTime = ct
    
    m,err:=getUsedMemory(ContainerId)
    if(err!=nil){
        return res,err
    }
    
    res.UsedMemory = m
    //return res
    return res,nil
} 

func (m ContainerMetrics) String() string {
	res, _ := json.Marshal(m)
    return (string(res))
}



func getCPUUtlizationPercent(ContainerID string) (float64,uint64,error) {    
    c,err:= config.ReadControlGroupFile(ContainerID,"cpuacct","cpuacct.usage")
    if(err != nil){
        //fmt.Println(err)
        return 0,0,err
    }
    conCpu,err := strconv.Atoi(strings.TrimSpace(c))
    if(err != nil){
        //fmt.Println(err)        
        return 0,0,err
    }
    fullCPU,err:=cpu.Times(false)
    if(err!=nil){
        //fmt.Println(err)        
        return 0,0,err
    }
    tbusy  := fullCPU[0].Total()-fullCPU[0].Idle
    //fmt.Println(busy)
    //fmt.Println(conCpu)
    cbusy:= float64(conCpu)/10000000.0
    return common.Round(cbusy/tbusy,0.5,3),uint64(cbusy),nil
}

//get memory in Megabytes
func getUsedMemory(containerID string) (int,error){
    v,err:= config.ReadControlGroupFile(containerID,"memory","memory.usage_in_bytes")
    if(err != nil){
        //fmt.Println(err)
        return 0,err
    }
    memory,err := strconv.Atoi(strings.TrimSpace(v))
    if(err != nil){
        //fmt.Println(err)        
        return 0,err
    }
    
    return memory/1048576,nil
}