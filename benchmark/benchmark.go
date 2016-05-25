package benchmark

import (
    "github.com/washraf/runcbm/common/conutil" 
    "github.com/washraf/runcbm/common/config"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

//Run ..
func Run(containerID string,n int) error { 
    condir,err:= config.FindContainerFolder(containerID)
    if(err!=nil){
        return err
    }
    measuresList :=make(Measures,0)
    for i := 1; i <= n; i++ {
        measure:= Measure{
        }
        u,err:=conutil.GetContainerUtilization(containerID) 
        if(err!=nil){
            return err
        }
        measure.ID = i
        measure.ProcessCount = 1 //to be updated    
        measure.MemorySize = u.UsedMemory
        command := exec.Command("time","-f","%e","runc","checkpoint",containerID)
        command.Dir = condir
        fmt.Println(command.Dir)
        r,err:=command.CombinedOutput()
        if(err!=nil){
            return err
        }
        measure.CheckpointTime,_ = strconv.ParseFloat(strings.TrimSpace(string(r)),64)	
        measure.Checkpointsize = 1234 
        command = exec.Command("time","-f","%e","runc","restore","-d",containerID)
        //command.Dir = "/containers/"+container+"/"
        command.Dir = condir
        
        r,err=command.CombinedOutput()
        if(err!=nil){
            return err
        }  
        measure.Restoretime,_ = strconv.ParseFloat(strings.TrimSpace(string(r)),64)	
        measuresList = append(measuresList,measure)    	
    }
    prinlist(measuresList)
    return nil
}

func prinlist(measuresList Measures){
    fmt.Printf("ID\tProcessCount\tMemorySize\tCheckpointTime\tCheckpointsize\tRestoretime\n")
    for _,m := range measuresList {
        fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", m.ID,m.ProcessCount,m.MemorySize,m.CheckpointTime,m.Checkpointsize,m.Restoretime)
    }
}