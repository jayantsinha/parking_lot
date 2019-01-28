package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	funcMapper := map[string]func([]string){
		CmdCreateParkingLot:         CreateParkingLot,
		CmdPark:                     ParkVehicle,
		CmdLeave:                    LeaveSlot,
		CmdStatus:                   GetStatus,
		CmdRegNumsForCarsWithColor:  GetRegNumsByColor,
		CmdSlotNumForRegNum:         GetSlotNumByRegNum,
		CmdSlotNumsForCarsWithColor: GetSlotsByColor,
		CmdExit:                     EndSession,
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Invalid command")
	}

	if len(flag.Args()) == 1 {
		ExecFile(flag.Args()[0])
	}

	if len(flag.Args()) > 0 {
		CallRespFunc(flag.Args(), funcMapper[flag.Args()[0]])
	}
}

//switch flag.Args()[0] {
//case CmdCreateParkingLot:
//	break;
//case CmdPark:
//	break;
//case CmdLeave:
//	break;
//case CmdStatus:
//	break;
//case CmdRegNumsForCarsWithColor:
//	break;
//case CmdSlotNumForRegNum:
//	break;
//case CmdSlotNumsForCarsWithColor:
//	break;
//case CmdExit:
//	break;
//default:
//	fmt.Println("Invalid command!")
//}