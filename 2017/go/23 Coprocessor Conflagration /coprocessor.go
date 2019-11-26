package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var registers = make(map[int]map[string]int)
var waiting = make(map[int]bool)
var queues = make(map[int][]int)
var count int = 0
var lastSound = 0
var tick = 0
var mulCount = 0

func main() {

	fmt.Println(runItAsGoCode())
}

func runItAsGoCode() int {
	b := 93      //set b 93
	c := b       //set c b
	b *= 100     //mul b 100
	b -= -100000 //sub b -100000
	c = b        //set c b
	c -= -17000  //sub c -17000
	f := 0
	//g := 0
	h := 0
	for true {
		f = 1 //    set f 1
		//d := 2 //    set d 2
		//e := 2 //        set e 2
		//g = 1 //to enter the loop
		if !big.NewInt(int64(b)).ProbablyPrime(0) {
			f = 0
		}
		// for d := 2; d < b; d++ {
		// 	//for g != 0 {
		// 	for e := 2; e < b; e++ {
		// 		//for g != 0 {
		// 		//g = d *e       //            set g d
		// 		//g *= e      //            mul g e
		// 		//g -= b      //            sub g b
		// 		if b == (d * e) { //                jnz g 2
		// 			f = 0 //                set f 0
		// 		}
		// 		//e++    //            sub e -1
		// 		//            set g e
		// 		//g = e - b //            sub g b
		// 	} //            jnz g -8
		// 	//d++    //        sub d -1
		// 	//g = d  //        set g d
		// 	//g -= b //        sub g b
		// } //        jnz g -13
		if f == 0 { //    jnz f 2
			h++ //    sub h -1
		}
		//g = b       //    set g b
		//g -= c      //    sub g c
		if c == b { //        jnz g 2
			return h //        jnz 1 3         # EXIT
		}
		b -= -17 //    sub b -17
	} //    jnz 1 -23
	return -100
}

func main_xx() {
	lines := fileToLines("input.txt")
	registers[0] = make(map[string]int)
	registers[1] = make(map[string]int)

	registers[0]["p"] = 0
	registers[1]["p"] = 1

	//Day 23 B
	registers[0]["a"] = 1

	registers[1]["done"] = 1
	for {
		tick++
		for pid := 0; pid < 1; pid++ {
			// if waiting[pid] {
			// 	registers[pid]["waiting"] = 1
			// } else {
			// 	registers[pid]["waiting"] = 0
			// }
			// registers[pid]["queueLength"] = len(queues[pid])
			// fmt.Println(registers[pid])
			// fmt.Println(lines[registers[pid]["pntr"]])

			if registers[pid]["done"] == 1 {
				continue
			}
			if waiting[pid] {
				continue
			}

			cmd, a, b := parseLine(pid, lines[registers[pid]["pntr"]])

			execCmd(pid, cmd, a, b)
			fmt.Printf("Line %d : %s (%d)\n", registers[pid]["pntr"]+1, lines[registers[pid]["pntr"]], registers[pid][a])

			registers[pid]["pntr"] = registers[pid]["pntr"] + 1
			if registers[pid]["pntr"] <= 0 || registers[pid]["pntr"] >= len(lines) {
				registers[pid]["done"] = 1
			}
		}
		//fmt.Println("------------------------------------------")
		if (waiting[0] || registers[0]["done"] == 1) && (waiting[1] || registers[1]["done"] == 1) {
			fmt.Println("Both terminated")
			break
		}

		if tick%1000000 == 0 {
			fmt.Print(".")
			//break
		}
		if tick%100000000 == 0 {
			fmt.Print(". \n")
			//break
		}

	}
	fmt.Printf("H: %d \n", registers[0]["h"])
	fmt.Printf("Mul count %d Tick:%d \n", mulCount, tick)
}

func execCmd(pid int, cmd string, a string, b int) {

	switch cmd {
	case "snd":
		break
	case "set":
		set(pid, a, b)
		break
	case "add":
		add(pid, a, b)
		break
	case "sub":
		sub(pid, a, b)
		break
	case "mul":
		mul(pid, a, b)
		break
	case "mod":
		mod(pid, a, b)
		break
	case "rcv":
		rcv(pid, a)
		break
	case "jgz":
		jgz(pid, a, b)
		break
	case "jnz":
		jnz(pid, a, b)
		break

	}
}
func parseLine(pid int, line string) (cmd string, a string, b int) {
	tokens := strings.Split(line, " ")

	x := 0
	if len(tokens) == 3 {
		ib, err := strconv.Atoi(tokens[2])
		if err != nil {
			ib = registers[pid][tokens[2]]
		}
		x = ib
	}
	return tokens[0], tokens[1], x
}

func snd(pid int, reg string) {

	val := registers[pid][reg]

	if pid == 0 {
		queues[1] = append(queues[1], val)
		waiting[1] = false
	} else {
		count++
		queues[0] = append(queues[0], val)
		waiting[0] = false
	}
}
func rcv(pid int, reg string) {
	if len(queues[pid]) == 0 {
		waiting[pid] = true
		registers[pid]["pntr"] = registers[pid]["pntr"] - 1 //So that we'll try this instruction again when the wait is over
		return
	}

	// if registers[pid][reg] != 0 {
	// 	if lastSound != 0 && pid == 0 {
	// 		fmt.Printf("Last sound for part one: %d \n", lastSound)
	// 	}
	// }
	registers[pid][reg] = queues[pid][0]
	queues[pid] = queues[pid][1:]

}
func set(pid int, reg string, val int) {
	registers[pid][reg] = val
}
func add(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] + val
}
func sub(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] - val
}
func mul(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] * val
	mulCount++
}
func mod(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] % val
}

func jgz(pid int, reg string, val int) {
	a, err := strconv.Atoi(reg)
	if err != nil {
		a = registers[pid][reg]
	}
	if a > 0 {
		registers[pid]["pntr"] = registers[pid]["pntr"] + (val - 1) //When I get out of here the pointer will ++, hence -1
	}
}
func jnz(pid int, reg string, val int) {
	a, err := strconv.Atoi(reg)
	if err != nil {
		a = registers[pid][reg]
	}
	if a != 0 {
		registers[pid]["pntr"] = registers[pid]["pntr"] + (val - 1) //When I get out of here the pointer will ++, hence -1
	}
}
