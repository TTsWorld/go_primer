package _interface

import (
	"fmt"
	"testing"
)

//go interface 实现多态

type USBer interface {
	start()
	stop()
}

type Computer struct{}

func (Computer) start() {
	fmt.Println("Computer started...")
}

func (Computer) stop() {
	fmt.Println("Computer stopped...")
}

type Phone struct{}

func (Phone) start() {
	fmt.Println("Phone started...")
}

func (Phone) stop() {
	fmt.Println("Phone stopped...")
}

func startMachine(ber USBer) {
	ber.start()
}

func stopMachine(ber USBer) {
	ber.stop()
}

func Test02(t *testing.T) {
	computer := Computer{}
	phone := Phone{}

	startMachine(computer)
	startMachine(phone)

	stopMachine(computer)
	stopMachine(phone)
}
