package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Trip contains all of the details for a single trip
type Trip struct {
	StartTime  time.Time
	EndTime    time.Time
	TotalHours float64
	Miles      float64
}

//NewTrip parses the record into a new trip
//It will error out of any parsing fails or the trip average MPH is below 5 or above 100
func NewTrip(record []string) (Trip, error) {
	if len(record) < 5 {
		return Trip{}, errors.New("Record is an invalid trip")
	}
	startTime, err := time.Parse("15:04", strings.TrimSpace(record[2]))
	if err != nil {
		return Trip{}, fmt.Errorf("Start time not in correct format : %v", strings.TrimSpace(record[2]))
	}
	endTime, err := time.Parse("15:04", strings.TrimSpace(record[3]))
	if err != nil {
		return Trip{}, fmt.Errorf("End time not in correct format : %v", strings.TrimSpace(record[3]))
	}
	miles, err := strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
	if err != nil {
		return Trip{}, fmt.Errorf("Miles not in correct format : %v", strings.TrimSpace(record[4]))
	}
	totalHours := endTime.Sub(startTime).Hours()

	avgMPH := averageMPH(miles, totalHours)
	if avgMPH < 5 || avgMPH > 100 {
		return Trip{}, fmt.Errorf("Trip is invalid because of average MPH : %v", avgMPH)
	}

	return Trip{
		StartTime:  startTime,
		EndTime:    endTime,
		TotalHours: totalHours,
		Miles:      miles,
	}, nil
}
