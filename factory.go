package main

import (
	"fmt"
	"strings"
)

func GetBareMetal(vendor string) (InterfaceBareMetal, error) {
	switch strings.ToUpper(vendor) {
	case "SUPERMICRO":
		return NewSupermicroDevice(), nil
	case "DELL":
		return NewDellDevice(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
