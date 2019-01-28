package main

import (
	"bufio"
	"fmt"
	"os"
)

func ExecFile(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		_ = fmt.Errorf("Error encountered while reading file: Err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func CallRespFunc(args []string, f func(args []string)) {
	f(args)
}

func CreateParkingLot(args []string) {
	fmt.Println(1)
}

func ParkVehicle(args []string) {
	fmt.Println(2)
}

func LeaveSlot(args []string) {
	fmt.Println(3)
}

func GetStatus(args []string) {
	fmt.Println(4)
}

func GetRegNumsByColor(args []string) {
	fmt.Println(5)
}

func GetSlotNumByRegNum(args []string) {
	fmt.Println(6)
}

func GetSlotsByColor(args []string) {
	fmt.Println(7)
}

func EndSession(args []string) {
	fmt.Println(8)
}