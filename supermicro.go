package main

type SupermicroDevice struct {
	BareMetal
}

func NewSupermicroDevice() InterfaceBareMetal {
	return &SupermicroDevice{
		BareMetal: BareMetal{
			name:        "AS-8125GS-TNMR2",
			vendor:      "supermicro",
			core:        0,
			memory:      0,
			diskPerGb:   0,
			biosVersion: "1.0.0",
		},
	}
}
