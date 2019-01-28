package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var singleCmds map[string]bool
var funcMapper map[string]func([]string) bool

func init() {
	flag.Parse()

	funcMapper = map[string]func([]string) bool{
		CmdCreateParkingLot:         CreateParkingLot,
		CmdPark:                     ParkVehicle,
		CmdLeave:                    LeaveSlot,
		CmdStatus:                   GetStatus,
		CmdRegNumsForCarsWithColor:  GetRegNumsByColor,
		CmdSlotNumForRegNum:         GetSlotNumByRegNum,
		CmdSlotNumsForCarsWithColor: GetSlotsByColor,
		CmdExit:                     EndSession,
		CmdVersion:                  ShowVersion,
	}
	singleCmds = map[string]bool{CmdStatus: true, CmdExit: true, CmdVersion: true}
}

func main() {
	if len(flag.Args()) < 1 {
		fmt.Println("Invalid command")
	}

	_, isSingleCmd := singleCmds[flag.Args()[0]]
	if len(flag.Args()) == 1  &&  !isSingleCmd {
		execFile(flag.Args()[0])
	}

	if len(flag.Args()) > 0 {
		_, isValidCmd := funcMapper[flag.Args()[0]]
		if !isValidCmd {
			fmt.Println("Unknown command")
			return
		}
		CallRespFunc(flag.Args(), funcMapper[flag.Args()[0]])
	}
}

// execFile executes commands from a text file
func execFile(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		_ = fmt.Errorf("Error encountered while reading file: Err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		args := separateFlags(line)
		processArgs(args)
	}
}

func separateFlags(line string) (args []string) {
	line = strings.TrimSpace(line)
	args = strings.Split(line, " ")
	return
}

func processArgs(args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid command")
	}

	_, isSingleCmd := singleCmds[args[0]]
	if len(args) == 1  &&  !isSingleCmd {
		execFile(args[0])
	}

	if len(args) > 0 {
		_, isValidCmd := funcMapper[args[0]]
		if !isValidCmd {
			fmt.Println("Unknown command")
			return
		}
		CallRespFunc(args, funcMapper[args[0]])
	}
}