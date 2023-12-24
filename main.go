package main

import "fmt"

func main() {
	supermicroDevice, _ := GetBareMetal("supermicro")
	printDetails(supermicroDevice)
	dellDevice, _ := GetBareMetal("dell")
	printDetails(dellDevice)
}

func printDetails(bm InterfaceBareMetal) {
	fmt.Printf("name : %s\n", bm.GetName())
	fmt.Printf("vendor : %s\n", bm.GetVendor())
}
