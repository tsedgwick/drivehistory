package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "", "specify file to read.")
}

func main() {
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	drivers := map[string]Driver{}
	err = process(file, drivers)
	if err != nil {
		log.Fatal(err)
	}

	err = write(os.Stdout, sortDrivers(drivers))
	if err != nil {
		log.Fatal(err)
	}
}

//write writes the drivers out in the format we are expecting
func write(writer io.Writer, drivers []Driver) error {

	for _, d := range drivers {
		output := fmt.Sprintf("%s: %d miles", d.Name, roundedTotalMiles(d.TotalMiles))
		if d.TotalMiles > 0 {
			output += fmt.Sprintf(" @ %d mph", averageMPH(d.TotalMiles, d.TotalHours))
		}

		fmt.Fprintf(writer, output+"\n")
	}

	return nil
}

func roundedTotalMiles(totalMiles float64) int {
	return int(math.Round(float64(totalMiles)))
}

func averageMPH(totalMiles, totalHours float64) int {
	if totalMiles == 0 {
		return 0
	}

	return int(math.Round(float64(totalMiles / totalHours)))
}
