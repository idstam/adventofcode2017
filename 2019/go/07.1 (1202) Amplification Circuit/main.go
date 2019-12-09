package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	lines := fileToLines("example.txt")
	strs := strings.Split(lines[0], ",")

	allPhases := IntPermutateArray([]int{5, 6, 7, 8, 9})

	lastOutput := 0

	maxThrust := 0
	for _, phases := range allPhases {

		lastOutput = 0

		vms := []VM1202{}
		for i, phase := range phases {
			vm := VM1202{
				Name: strconv.Itoa(i) + ":" + strconv.Itoa(phase),
				mem:  StringToIntArray(strs),
			}
			vm.inputs = make(chan int, 100)
			vm.output = make(chan int)
			vm.outputMode = "Console"
			vm.inputMode = "Channel"
			vm.logLevel = 99
			//vm.inputs <- phase

			vms = append(vms, vm)
		}
		for vms[4].state != "Done" {
			vms[0].inputs <- phases[0]
			vms[0].inputs <- vms[4].lastOutput

			vms[0].Run()

			vms[1].inputs <- phases[1]
			vms[1].inputs <- vms[0].lastOutput
			vms[1].Run()

			vms[2].inputs <- phases[2]
			vms[2].inputs <- vms[1].lastOutput
			vms[2].Run()
			vms[3].inputs <- phases[3]
			vms[3].inputs <- vms[2].lastOutput
			vms[3].Run()
			vms[4].inputs <- phases[4]
			vms[4].inputs <- vms[3].lastOutput
			vms[4].Run()
		}
		lastOutput = vms[4].lastOutput
		//fmt.Println()

		//vms[0].inputs <- 0

		//		wg := sync.WaitGroup{}
		//		wg.Add(len(phases) - 1)
		//out := make(chan int)

		// for _, vm := range vms {
		// 	go func(vm VM1202) {
		// 		vm.Run()
		// 		wg.Done()
		// 	}(vm)
		// }
		//		wg.Wait()
		//result := <-out
		// fmt.Println(result)
		// //lastOutput = <-vms[4].output
		// fmt.Println("More thrust ", result, phases)

		if lastOutput > maxThrust {
			fmt.Println("More thrust ", lastOutput, phases)
			//More thrust  206580 [2 0 1 4 3]
			//98765 too low
			maxThrust = lastOutput
		}
	}

}
