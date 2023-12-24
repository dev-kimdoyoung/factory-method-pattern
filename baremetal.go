package main

type BareMetal struct {
	name        string
	vendor      string
	core        int
	memory      int
	diskPerGb   int
	biosVersion string
}

func (b *BareMetal) SetName(name string) {
	b.name = name
}

func (b *BareMetal) SetVendor(vendor string) {
	b.vendor = vendor
}

func (b *BareMetal) SetCore(core int) {
	b.core = core
}

func (b *BareMetal) SetMemory(memory int) {
	b.memory = memory
}

func (b *BareMetal) SetDiskPerGb(disk int) {
	b.diskPerGb = disk
}

func (b *BareMetal) SetBiosVersion(version string) {
	b.biosVersion = version
}

func (b *BareMetal) GetName() string {
	return b.name
}

func (b *BareMetal) GetVendor() string {
	return b.vendor
}

func (b *BareMetal) GetCore() int {
	return b.core
}

func (b *BareMetal) GetMemory() int {
	return b.memory
}

func (b *BareMetal) GetDiskPerGb() int {
	return b.diskPerGb
}

func (b *BareMetal) GetBiosVersion() string {
	return b.biosVersion
}
