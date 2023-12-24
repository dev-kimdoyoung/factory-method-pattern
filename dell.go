package main

type DellDevice struct {
	BareMetal
}

func NewDellDevice() InterfaceBareMetal {
	return &DellDevice{
		BareMetal: BareMetal{
			name:        "R350",
			vendor:      "dell",
			core:        0,
			memory:      0,
			diskPerGb:   0,
			biosVersion: "1.0.0",
		},
	}
}
