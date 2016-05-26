# runcbm
Runc Continer tool Bench Marking tool

This is a tool to find speed of container checkpoint and restore 

Measure Container Performace
[x] Measure No of Processes
[x] Measure Memory size
[x] Measure Cpu

Check Point and Restoration Bench Mark
[x] Measure Checkpoint time
[x] Measure Resotration time
[x] Measure Checkpoint size
[x] Configure Repeat Test
[] Find Averages and Maximum

# BechMarking Example
```
NAME:
   runcbm bm - benchmark a container 

USAGE:
   runcbm bm [command options] COMMAND [arguments...]

OPTIONS:
   --id 	Container ID
   -n "5"	Number of Trails
   --dir 	location of the bundle folder
   
flag provided but not defined: -bundle
walid@ubuntu:~$ sudo ./runcbm bm --id redis --dir /containers/redis --n 5
Bench Mark Container ID redis 
 Starting Container
ID	ProcessCount	MemorySize	CheckpointTime	Checkpointsize	Restoretime
1	1	0	0.2	1	0.4
2	1	0	0.19	1	0.4
3	1	0	0.18	1	0.41
4	1	0	0.21	1	0.4
5	1	0	0.19	1	0.4
Kill Container
Delete Container
```

# Performace Example

```
walid@ubuntu:~/GoCode/src/github.com/washraf/runcbm$ sudo ./runcbm measure -h
NAME:
   runcbm measure - measure container load

USAGE:
   runcbm measure [command options] COMMAND [arguments...]

OPTIONS:
   --id 	Container ID
   
walid@ubuntu:~/$ sudo ./runcbm measure --id myubuntutest
Measure Container ID myubuntutest 
 {"CPUTime":850,"pCount":2,"usedMemory(MB)":0,"rootSize(MB)":213}
```