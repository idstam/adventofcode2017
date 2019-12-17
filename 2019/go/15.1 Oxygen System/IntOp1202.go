package main

import (
	"fmt"
	"log"
)

type VM1202 struct {
	Name           string
	mem            map[int64]int64
	Output         chan int64
	LastOutput     int64
	Input          chan int64
	InputValue     int64
	InputFunction  func(*VM1202) int64
	InputMode      string
	OutputMode     string
	OutputFunction func(*VM1202, int64)
	LogLevel       int
	LogMode        string
	LogLines       []string
	relativeBase   int64
	State          string
}

func (vm *VM1202) Run() {
	vm.Log(1, vm.Name, "RUN")
	vm.State = "Running"
	ptr := int64(0)
	for ptr >= 0 {
		ptr = vm.Exec(ptr)
	}
	vm.State = "Done"
}
func (vm *VM1202) Exec(ptr int64) int64 {

	//fmt.Println(vm.mem)
	fullOp := vm.mem[ptr]
	ma, mb, mc, op := vm.ParseOpCode(fullOp)
	if op == 99 {
		//fmt.Println("Exit")
		return -1
	}
	switch op {
	case 1:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		v3 := vm.GetValue(ptr+3, 1)
		dest := v3
		if ma == 2 {
			dest += vm.relativeBase
		}

		vm.mem[dest] = v1 + v2
		ptr += 4
	case 2:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		v3 := vm.GetValue(ptr+3, 1)
		dest := v3
		if ma == 2 {
			dest += vm.relativeBase
		}

		vm.mem[dest] = v1 * v2
		ptr += 4
	case 3:
		v1 := vm.GetValue(ptr+1, 1)
		dest := v1
		if mc == 2 {
			dest += vm.relativeBase
		}
		vm.mem[dest] = vm.GetInput()
		ptr += 2
	case 4:
		v1 := vm.GetValue(ptr+1, mc)
		val := v1
		vm.SendOutput(val)
		ptr += 2
	case 5:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		if v1 != 0 {
			ptr = v2
		} else {
			ptr += 3
		}
	case 6:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		if v1 == 0 {
			ptr = v2
		} else {
			ptr += 3
		}
	case 7:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		v3 := vm.GetValue(ptr+3, 1)

		dest := v3
		if ma == 2 {
			dest += vm.relativeBase
		}

		if v1 < v2 {
			vm.mem[dest] = 1
		} else {
			vm.mem[dest] = 0
		}
		ptr += 4
	case 8:
		v1 := vm.GetValue(ptr+1, mc)
		v2 := vm.GetValue(ptr+2, mb)
		v3 := vm.GetValue(ptr+3, 1)
		dest := v3
		if ma == 2 {
			dest += vm.relativeBase
		}

		if v1 == v2 {
			vm.mem[dest] = 1
		} else {
			vm.mem[dest] = 0
		}
		ptr += 4

	case 9:
		v1 := vm.GetValue(ptr+1, mc)
		val := v1
		vm.relativeBase += val
		ptr += 2
	case 99:
		vm.Log(1, "Exit")
		return -1
	default:
		log.Fatalf("Unknown OP Code %d \n", op)

	}
	//fmt.Println("After 225:", mem[225])
	return ptr
}
func (vm *VM1202) OpName(op int64) string {
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
	case 9:
		return "Change relative base with a"
	case 99:
		return "Exit"
	default:
		return fmt.Sprintf("UNKNOWN %d", op)

	}
}

func (vm *VM1202) GetValue(adress int64, mode int64) int64 {
	ret := int64(0)
	from := int64(0)

	switch mode {
	case 0:
		from = vm.mem[adress]
		if from >= 0 {
			ret = vm.mem[from]
		}
	case 1:
		ret = vm.mem[adress]
	case 2:
		from = vm.mem[adress] + vm.relativeBase
		if from >= 0 {
			ret = vm.mem[from]
		}
	}
	return ret
}
func (vm *VM1202) GetInput() int64 {
	vm.Log(0, vm.Name, "GetInput", "in")
	ret := int64(0)

	switch vm.InputMode {
	case "Console":
		fmt.Scanf("%d", &ret)
	case "Channel":
		ret = <-vm.Input
	case "Value":
		ret = vm.InputValue
	case "Function":
		ret = vm.InputFunction(vm)
	default:
		log.Panic("UNKNOWN INPUT MODE " + vm.InputMode)
	}

	vm.Log(0, vm.Name, "GetInput", "out", ret)
	return ret
}
func (vm *VM1202) SendOutput(val int64) {
	vm.Log(0, vm.Name, "SendOutput", vm.OutputMode, "in", val)
	vm.LastOutput = val
	switch vm.OutputMode {
	case "Console":
		fmt.Println(vm.Name, "SendOutput", "->", val)
	case "Channel":
		vm.Output <- val
	case "Function":
		vm.OutputFunction(vm, val)
	case "Silent":
	default:
		log.Panic("UNKNOWN OUTPUT MODE " + vm.OutputMode)
	}
	vm.Log(0, vm.Name, "SendOutput", vm.OutputMode, "out")

}
func (vm *VM1202) ParseOpCode(in int64) (int64, int64, int64, int64) {
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
		vm.LogLines = append(vm.LogLines, fmt.Sprint(message...))
		fmt.Println(message...)
	}
}
func (vm *VM1202) DumpLogLines() {
	for _, l := range vm.LogLines {
		fmt.Println(l)
	}
}
func (vm *VM1202) DumpInfo() {
	fmt.Println("State", vm.State)
	fmt.Println("LastOutput", vm.LastOutput)
}
