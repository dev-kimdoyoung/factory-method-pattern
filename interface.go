package main

type InterfaceBareMetal interface {
	SetName(string)
	SetVendor(string)
	SetCore(int)
	SetMemory(int)
	SetDiskPerGb(int)
	SetBiosVersion(string)
	GetName() string
	GetVendor() string
	GetCore() int
	GetMemory() int
	GetDiskPerGb() int
	GetBiosVersion() string
}
