package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile     = "input1.txt"
)

type instruction struct {
	operation string
	arguments [2]string
}

var inputCounter int
var expression [4]*string

func main() {
	solution1(inputFile)
}

func solution1(fileName string) {
	instructions := readInput(fileName)
	//fmt.Println(instructions)
	for i := 0; i < 4; i++ {
		empty := "0"
		expression[i] = &empty
	}
	for _, instr := range instructions {
		//if inputCounter > 2 {
		//	break
		//}
		update(instr)
		//fmt.Println(fmt.Sprintf("%s: x: %s", instr, *expression[1]))
	}
	printExpressions()
}

func get(in string) string {
	switch in {
	case "w":
		return *expression[0]
	case "x":
		return *expression[1]
	case "y":
		return *expression[2]
	case "z":
		return *expression[3]
	default:
		return in
	}
}

func write(state, value string) {
	switch state {
	case "w":
		*expression[0] = value
	case "x":
		*expression[1] = value
	case "y":
		*expression[2] = value
	case "z":
		*expression[3] = value
	}
}

func printExpressions() {
	for i := 3; i < 4; i++ {
		fmt.Print(fmt.Sprintf("%s", *expression[i]))
		fmt.Println()
	}
}

func update(instr instruction) {
	switch instr.operation {
	case "inp":
		write(instr.arguments[0], fmt.Sprintf("in[%d]", inputCounter))
		inputCounter++
		printExpressions()
		write("z", fmt.Sprintf("z%d", inputCounter-1))
		//write("x", "0")
		//write("y", "0")
		//write("z", "0")
	case "add":
		base := get(instr.arguments[0]) + "+"
		if base == "0+" {
			base = ""
		}
		if _, err := strconv.Atoi(instr.arguments[1]); err == nil {
			write(instr.arguments[0], fmt.Sprintf("%s%s", base, instr.arguments[1]))
		} else {
			write(instr.arguments[0], fmt.Sprintf("%s%s", base, get(instr.arguments[1])))
		}
	case "mul":
		if get(instr.arguments[0]) == "0" {
			return
		}
		if get(instr.arguments[1]) == "1" {
			return
		}
		if instr.arguments[1] == "0" {
			write(instr.arguments[0], "0")
			return
		}
		if _, err := strconv.Atoi(instr.arguments[1]); err == nil {
			write(instr.arguments[0], fmt.Sprintf("(%s)*%s", get(instr.arguments[0]), instr.arguments[1]))
		} else {
			if get(instr.arguments[1]) == "0" {
				write(instr.arguments[0], "0")
				return
			}
			write(instr.arguments[0], fmt.Sprintf("(%s)*(%s)", get(instr.arguments[0]), get(instr.arguments[1])))
		}
	case "div":
		if get(instr.arguments[0]) == "0" {
			return
		}
		if get(instr.arguments[1]) == "1" {
			return
		}
		if _, err := strconv.Atoi(instr.arguments[1]); err == nil {
			write(instr.arguments[0], fmt.Sprintf("(%s)/%s", get(instr.arguments[0]), instr.arguments[1]))
		} else {
			write(instr.arguments[0], fmt.Sprintf("(%s)/(%s)", get(instr.arguments[0]), get(instr.arguments[1])))
		}
	case "mod":
		if get(instr.arguments[0]) == "0" {
			return
		}
		if _, err := strconv.Atoi(instr.arguments[1]); err == nil {
			write(instr.arguments[0], "("+get(instr.arguments[0])+")%"+instr.arguments[1])
		} else {
			write(instr.arguments[0], "("+get(instr.arguments[0])+")%("+get(instr.arguments[1])+")")
		}
	case "eql":
		//if len(get(instr.arguments[0])) > 1 && strings.HasPrefix(get(instr.arguments[1]), "in") {
		//	write(instr.arguments[0], "0")
		//	return
		//}
		if get(instr.arguments[0]) == "0" && strings.HasPrefix(get(instr.arguments[1]), "in") {
			write(instr.arguments[0], "0")
			return
		}
		if get(instr.arguments[0]) == get(instr.arguments[1]) {
			write(instr.arguments[0], "1")
			return
		}
		if _, err := strconv.Atoi(instr.arguments[1]); err == nil {
			write(instr.arguments[0], fmt.Sprintf("(%s)==%s", get(instr.arguments[0]), instr.arguments[1]))
		} else {
			write(instr.arguments[0], fmt.Sprintf("(%s)==(%s)", get(instr.arguments[0]), get(instr.arguments[1])))
		}
	}
}

func readInput(fileName string) []instruction {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	instructions := make([]instruction, len(rows))
	for i, row := range rows {
		spl := strings.Split(row, " ")
		op := spl[0]
		var args [2]string
		if op == "inp" {
			args = [2]string{spl[1], ""}
		} else {
			args = [2]string{spl[1], spl[2]}
		}
		instructions[i] = instruction{
			operation: op,
			arguments: args,
		}
	}
	return instructions
}
