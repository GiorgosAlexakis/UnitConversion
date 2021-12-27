package main

import (
	"Looper2074/UnitConversion/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	typeOfUnit := os.Args[1]

	for _, arg := range os.Args[2:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		if typeOfUnit == "temperature" {
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	}
}
