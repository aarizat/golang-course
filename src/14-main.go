package main

import "fmt"


type pc struct {
	ram int
	disk int
	brand string
}

func (PC pc) ping() {
	fmt.Println(PC.brand, "Pong")
}

func (PC *pc) doubleRam() {
	PC.ram *= 2
}

func (PC pc) String() string {
	return fmt.Sprintf("It has %d GB RAM, %d GB Disk and it is a %s", PC.ram, PC.disk, PC.brand)
}

func main() {
	a := 50
	b := &a
	fmt.Println(b)
	*b = 1000
	fmt.Println(a)

	// define pc
	myPC := pc{ram: 8, disk: 500, brand: "Lenovo"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.doubleRam()
	fmt.Println(myPC)

	myPC.doubleRam()
	fmt.Println(myPC)

	// customize ouput
	fmt.Println(myPC)
}