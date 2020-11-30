package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")

	allPhases := IntPermutateArray([]int{5, 6, 7, 8, 9})

	lastOutput := 0

	maxThrust := -1
	for _, phases := range allPhases {

		lastOutput = 0

		vms := []VM1202{}
		for i, phase := range phases {
			vm := VM1202{
				Name: strconv.Itoa(i) + ":" + strconv.Itoa(phase),
				mem:  StringToIntArray(strs),
			}
			vm.Input = make(chan int, 2)
			vm.Output = make(chan int)
			vm.OutputMode = "Channel"
			vm.InputMode = "Channel"
			vm.LogLevel = 99
			//vm.inputs <- phase

			vms = append(vms, vm)
		}

		vms[0].Input <- phases[0]
		vms[1].Input <- phases[1]
		vms[2].Input <- phases[2]
		vms[3].Input <- phases[3]
		vms[4].Input <- phases[4]

		vms[0].Input <- 0

		go vms[0].Run()
		go vms[1].Run()
		go vms[2].Run()
		go vms[3].Run()
		go vms[4].Run()

		for vms[4].State != "Done" {
			//time.Sleep( 0* time.Second)

			select {
			case val0 := <-vms[0].Output:

				vms[1].Input <- val0

			case val1 := <-vms[1].Output:

				vms[2].Input <- val1

			case val2 := <-vms[2].Output:

				vms[3].Input <- val2

			case val3 := <-vms[3].Output:

				vms[4].Input <- val3

			case val4 := <-vms[4].Output:

				vms[0].Input <- val4
			default:
				//fmt.Println(vms[4].State, "B")

			}

		}
		lastOutput = vms[4].LastOutput
		if lastOutput > maxThrust {
			fmt.Println("More thrust ", lastOutput, phases)
			//More thrust  206580 [2 0 1 4 3]
			//98765 too low
			maxThrust = lastOutput
		}
	}

	fmt.Println(maxThrust)
}
