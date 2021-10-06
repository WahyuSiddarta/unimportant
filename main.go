package main

import (
	"fmt"
	"runtime"
	"time"

	//"time"

	"github.com/ruangnyaman/rna-ecommerce-backend/api"
	"github.com/shirou/gopsutil/cpu"
)

func main() {

	runtime.GOMAXPROCS(4 * runtime.NumCPU())
	{
		fmt.Println("VCPU Proc :", runtime.NumCPU())
	}
	Ticktoc := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-Ticktoc.C:
				go func() {
					usage, _ := cpu.Percent(5*time.Second, false)
					fmt.Printf("########## CPU %3.2f %% ###### routine:%d\n", usage[0], runtime.NumGoroutine())
				}()
			}
		}
	}()

	api.Init().SetupRoutes()
}
