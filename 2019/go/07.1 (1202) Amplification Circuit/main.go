package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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
				Name: strconv.Itoa(phase),
				mem:  StringToIntArray(strs),
			}
			vm.inputs = make(chan int, 100)
			vm.output = make(chan int, 100)
			vm.inputs <- phase

			if i < len(vms)-1 {

				vms[i+1].inputs = vm.output
			}

			vms = append(vms, vm)
		}

		wg := sync.WaitGroup{}
		wg.Add(len(phases))
		out := make(chan int)

		for _, vm := range vms {
			go func(vm VM1202) {
				vm.Run()
				wg.Done()
			}(vm)
		}
		wg.Wait()
		result := <-out
		fmt.Println(result)
		//lastOutput = <-vms[4].output
		fmt.Println("More thrust ", result, phases)

		if lastOutput > maxThrust {
			fmt.Println("More thrust ", lastOutput, phases)
			//More thrust  206580 [2 0 1 4 3]
			maxThrust = lastOutput
		}
	}

}
