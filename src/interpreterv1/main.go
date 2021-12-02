package main

import (
	"crank/repl"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

type SysInfo struct {
	Platform string `bson:platform`
	CPU      string `bson:cpu`
}

func main() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()

	info := new(SysInfo)

	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName

	fmt.Print(`
	=============================================================================

	/$$$$$$                               /$$      
	/$$__  $$                             | $$      
   | $$  \__/  /$$$$$$  /$$$$$$  /$$$$$$$ | $$   /$$
   | $$       /$$__  $$|____  $$| $$__  $$| $$  /$$/
   | $$      | $$  \__/ /$$$$$$$| $$  \ $$| $$$$$$/ 
   | $$    $$| $$      /$$__  $$| $$  | $$| $$_  $$ 
   |  $$$$$$/| $$     |  $$$$$$$| $$  | $$| $$ \  $$
	\______/ |__/      \_______/|__/  |__/|__/  \__/
	=============================================================================
												   
	`)

	fmt.Println("Running on: " + info.Platform + " CPU model: " + info.CPU)

	repl.Start(os.Stdin, os.Stdout)
}
