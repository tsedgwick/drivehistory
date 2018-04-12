package main

import (
	"sort"
)

//Driver is a struct that contains trip details the driver has taken
type Driver struct {
	Name       string
	Trips      []Trip
	TotalMiles float64
	TotalHours float64
}

//Drivers is a named slice used for sorting
type Drivers []Driver

//sortDrivers takes in a map and returns back the sorted values
func sortDrivers(drivers map[string]Driver) []Driver {
	sortedDrivers := make(Drivers, len(drivers))
	var count int
	for _, driver := range drivers {
		sortedDrivers[count] = driver
		count++
	}

	sort.Sort(sortedDrivers)
	return sortedDrivers
}

func (d Drivers) Len() int {
	return len(d)
}

func (d Drivers) Less(i, j int) bool {
	if d[i].TotalMiles != d[j].TotalMiles {
		return d[i].TotalMiles > d[j].TotalMiles
	}
	return false
}

func (d Drivers) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
