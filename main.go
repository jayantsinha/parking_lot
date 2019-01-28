package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	funcMapper := map[string]func([]string) bool{
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
	singleCmds := map[string]bool{CmdStatus: true, CmdExit: true, CmdVersion: true}

	if len(flag.Args()) < 1 {
		fmt.Println("Invalid command")
	}

	_, isSingleCmd := singleCmds[flag.Args()[0]]
	if len(flag.Args()) == 1  &&  !isSingleCmd {
		ExecFile(flag.Args()[0])
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
