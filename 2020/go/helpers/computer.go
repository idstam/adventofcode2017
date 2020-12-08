package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

//Computer is the computer
type Computer struct {
	PC           int //Program Counter
	Accumulator  int //Accumulator
	Instructions []Instruction
	Mode         string
}

//Instruction is a parsed line of code
type Instruction struct {
	Code    string
	Name    string
	IntValA int
}

//Reset clear all registers
func (c *Computer) Reset() {
	c.PC = 0
	c.Accumulator = 0
}

//Load read and parse the code
func (c *Computer) Load(code []string) {
	c.Instructions = c.MapCodeToInstruction(code)
}

//Run execute from the current program counter
func (c *Computer) Run() {
	lastInstruction := 0
	calledInstructions := map[int]bool{}
	for c.PC < len(c.Instructions) && c.PC >= 0 {
		lastInstruction = c.PC

		if calledInstructions[c.PC] {
			fmt.Println("Acumulator when instruction run twice", c.Accumulator)
			if c.Mode == "08" {
				break
			}
		}

		calledInstructions[c.PC] = true
		fmt.Printf("Pre exec PC=%d Instr=%+v \n", c.PC, c.Instructions[c.PC])
		c.PC = c.Exec()
		fmt.Printf("Post exec PC=%d Accumulator=%d \n", c.PC, c.Accumulator)

	}
	if c.PC < 0 || c.PC >= len(c.Instructions) {
		fmt.Println("Program counter is out of range", c.PC, lastInstruction)
	}

}

//Boot initiates the computer
func (c *Computer) Boot(code []string) {
	c.Reset()
	c.Load(code)
	c.Run()

}

//MapCodeToInstruction Parse a line of code
func (c *Computer) MapCodeToInstruction(code []string) []Instruction {

	ret := []Instruction{}
	for _, l := range code {
		tokens := strings.Split(l, " ")
		intValA, _ := strconv.Atoi(tokens[1])
		ret = append(ret, Instruction{Code: l, Name: tokens[0], IntValA: intValA})
	}
	return ret
}

//Exec Execute the current line of code
func (c *Computer) Exec() int {
	instr := c.Instructions[c.PC]

	switch instr.Name {
	case "acc":
		return c.Acc(instr.IntValA)
	case "jmp":
		return c.Jmp(instr.IntValA)
	case "nop":
		return c.Nop(instr.IntValA)

	}
	fmt.Printf("Unknown operation %+v \n", instr)
	return -1
}

//Acc - Increment or decrement accumulator
func (c *Computer) Acc(val int) int {
	c.Accumulator += val
	return c.PC + 1
}

//Jmp - Move program counter
func (c *Computer) Jmp(val int) int {
	return c.PC + val
}

//Nop - Increment PC and do nothing else
func (c *Computer) Nop(val int) int {
	return c.PC + 1
}
