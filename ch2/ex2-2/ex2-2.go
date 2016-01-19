// Ex2-2 performs temp, len, and wt conversions on its args or Stdin.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Derek701/gopl-solutions/ch2/ex2-1/tempconv"
	"github.com/Derek701/gopl-solutions/ch2/ex2-2/lenconv"
	"github.com/Derek701/gopl-solutions/ch2/ex2-2/wtconv"
)

func main() {
	if len(os.Args) > 1 { // Display conversions for any arguments
		for _, arg := range os.Args[1:] {
			display(arg)
		}
	} else { // Otherwise scan and display conversions from Stdin
		for {
			var input string
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %v\n", err)
				os.Exit(1)
			}
			display(input)
		}
	}
}

func display(input string) {
	fmt.Println()
	t, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
	ft := lenconv.Feet(t)
	m := lenconv.Meters(t)
	fmt.Printf("%s = %s, %s = %s\n",
		ft, lenconv.FToM(ft), m, lenconv.MToF(m))
	lb := wtconv.Pounds(t)
	kg := wtconv.Kilograms(t)
	fmt.Printf("%s = %s, %s = %s\n",
		lb, wtconv.LbToKg(lb), kg, wtconv.KgToLb(kg))
}
