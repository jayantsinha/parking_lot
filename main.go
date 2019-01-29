package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// singleCmds holds all commands which does not have any argument
var singleCmds map[string]bool
// funcMapper is a map of command to its handler
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
	if len(flag.Args()) == 1 {
		_, isSingleCmd := singleCmds[flag.Args()[0]]
		if !isSingleCmd {
			execFile(flag.Args()[0])
			os.Exit(1)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Welcome to Parking Lot Manager")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		cmd := separateFlags(line)
		command := cmd[0]
		if cmd[0] != CmdCreateParkingLot {
			if cmd[0] == CmdExit {
				EndSession(nil)
			}
			fmt.Println("Please create a parking lot first using `create_parking_lot <numSlots>`")
		}
		if len(cmd) != 2 {
			fmt.Println("Invalid syntax")
		}
		ret := CreateParkingLot(cmd)
		if !ret {
			fmt.Println("Invalid input; unable to create parking lot.")
		}
		for command != CmdExit {
			line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)
			cmd = separateFlags(line)
			if len(cmd) == 0 {
				fmt.Println("Invalid command")
				continue
			}
			if len(cmd) > 0 {
				_, isValidCmd := funcMapper[cmd[0]]
				if !isValidCmd {
					fmt.Println("Unknown command")
					continue
				}
				CallRespFunc(cmd, funcMapper[cmd[0]])
			}
		}
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

// separateFlags separates the commands and arguments from the given line
func separateFlags(line string) (args []string) {
	line = strings.TrimSpace(line)
	args = strings.Split(line, " ")
	return
}

// processArgs processes the commands with the arguments
func processArgs(args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid command")
	}

	_, isSingleCmd := singleCmds[args[0]]
	if len(args) == 1 && !isSingleCmd {
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
