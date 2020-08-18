package main

import (
	"fmt"
	"log"
)

type VM1202 struct {
	Name       string
	mem        []int
	Output     chan int
	LastOutput int
	Input      chan int
	InputMode  string
	OutputMode string
	LogLevel   int

	State string
	ptr   int
}

func (vm *VM1202) Run() {
	vm.Log(1, vm.Name, "RUN")
	vm.State = "Running"
	ptr := 0
	vm.state = "Running"
	for ptr >= 0 {
		ptr = vm.Exec(ptr)
	}
	vm.State = "Done"
}
func (vm *VM1202) Exec(ptr int) int {
	fullOp := vm.mem[ptr]
	_, mb, mc, op := vm.ParseOpCode(fullOp)
	if op == 99 {
		if vm.logLevel == 0 {
			fmt.Println(vm.Name, "Exit")
		}
		return -1
	}
	v1, v2, v3 := vm.GetValues(ptr)
	//fmt.Println(OpName(op), ": ", v1,mc, v2,mb, v3,ma)
	//fmt.Println("Before 225:", mem[225])

	switch op {
	case 1:
		dest := v3[1]
		vm.ExpandTape(dest + 1)
		vm.mem[dest] = v1[mc] + v2[mb]
		ptr += 4
	case 2:
		dest := v3[1]
		vm.ExpandTape(dest + 1)
		vm.mem[dest] = v1[mc] * v2[mb]
		ptr += 4
	case 3:
		dest := v1[1]
		vm.ExpandTape(dest + 1)
		vm.mem[dest] = vm.GetInput()
		ptr += 2
	case 4:
		val := v1[mc]
		vm.SendOutput(val)
		ptr += 2
	case 5:
		if v1[mc] != 0 {
			ptr = v2[mb]
		} else {
			ptr += 3
		}
	case 6:
		if v1[mc] == 0 {
			ptr = v2[mb]
		} else {
			ptr += 3
		}
	case 7:
		dest := v3[1]
		vm.ExpandTape(dest + 1)

		if v1[mc] < v2[mb] {
			vm.mem[dest] = 1
		} else {
			vm.mem[dest] = 0
		}
		vm.ptr += 4
	case 8:
		dest := v3[1]
		vm.ExpandTape(dest + 1)

		if v1[mc] == v2[mb] {
			vm.mem[dest] = 1
		} else {
			vm.mem[dest] = 0
		}
		vm.ptr += 4

	case 99:
		vm.Log(1, "Exit")
		return -1
	default:
		log.Fatalf("Unknown OP Code %d \n", op)

	}
	//fmt.Println("After 225:", mem[225])
	return ptr
}
func (vm *VM1202) OpName(op int) string {
	switch op {
	case 1:
		return "Add a+b -> c"
	case 2:
		return "Add a*b -> c"
	case 3:
		return "Input -> a"
	case 4:
		return "Output a"
	case 5:
		return "JIT a=!0 => b"
	case 6:
		return "JIF a==0 => b"
	case 7:
		return "LT a < b -> c"
	case 8:
		return "EQ a == b -> c"
	case 99:
		return "Exit"
	default:
		return fmt.Sprintf("UNKNOWN %d", op)

	}
}
func (vm *VM1202) ExpandTape(newSize int) {
	if len(vm.mem) >= newSize {
		return
	}

	for i := len(vm.mem); i <= newSize; i++ {
		vm.mem = append(vm.mem, 0)
	}
}
func (vm *VM1202) GetValues(adress int) ([]int, []int, []int) {
	v1 := []int{0, 0}
	v2 := []int{0, 0}
	v3 := []int{0, 0}
	v1 = vm.GetValue(adress + 1)
	v2 = vm.GetValue(adress + 2)
	v3 = vm.GetValue(adress + 3)
	return v1, v2, v3
}
func (vm *VM1202) GetValue(adress int) []int {
	vm.ExpandTape(IntMax(adress, vm.mem[adress]) + 1)
	ret := -9999999
	from := vm.mem[adress]
	if from >= 0 {
		ret = vm.mem[from]
	}
	return []int{ret, vm.mem[adress]}

}
func (vm *VM1202) GetInput() int {
	vm.Log(0, vm.Name, "GetInput", "in")
	ret := 0

	switch vm.InputMode {
	case "Console":
		fmt.Scan(&ret)
	case "Channel":
		ret = <-vm.Input
	default:
		log.Panic("UNKNOWN INPUT MODE " + vm.InputMode)
	}

	vm.Log(0, vm.Name, "GetInput", "out", ret)
	return ret
}
func (vm *VM1202) SendOutput(val int) {
	vm.Log(0, vm.Name, "SendOutput", vm.OutputMode, "in", val)
	vm.LastOutput = val
	switch vm.OutputMode {
	case "Console":
		fmt.Println(vm.Name, "SendOutput", "->", val)
	case "Channel":
		vm.Output <- val
	case "Silent":
	default:
		log.Panic("UNKNOWN OUTPUT MODE " + vm.OutputMode)
	}
	vm.Log(0, vm.Name, "SendOutput", vm.OutputMode, "out")

}
func (vm *VM1202) ParseOpCode(in int) (int, int, int, int) {
	op := in % 100
	in -= op
	c := (in % 1000) / 100
	in -= c * 100
	b := (in % 10000) / 1000
	in -= b * 1000
	a := (in % 100000) / 10000
	return a, b, c, op
}
func (vm *VM1202) Log(level int, message ...interface{}) {
	if level >= vm.LogLevel {
		fmt.Println(message)
	}
}
