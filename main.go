package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	data := `
	Driver Dan
	Driver Alex
	Driver Bob
	Trip Dan 07:15 07:45 17.3
	Trip Dan 06:12 06:32 21.8
	Trip Alex 12:01 13:16 42.0
	Invalid
`

	drivers := map[string]Driver{}

	for rec := range processCSV(strings.NewReader(data)) {
		if rec.err != nil {
			fmt.Printf("Line : %v : has an error : %v \n", rec.record, rec.err)
			continue
		}
		switch len(rec.record) {
		case 2:
			name := strings.TrimSpace(rec.record[1])
			driver, ok := drivers[name]
			if !ok {
				driver = Driver{
					Name:  name,
					Trips: make([]Trip, 0),
				}
			}

			drivers[driver.Name] = driver
		case 5:
			name := strings.TrimSpace(rec.record[1])
			driver, ok := drivers[name]

			if !ok {
				fmt.Printf("Trip is for unregistered driver : %v\n", rec.record)
				continue
			}
			trip, err := NewTrip(rec.record)
			if err != nil {
				fmt.Println(err)
			}

			driver.Trips = append(driver.Trips, trip)
			driver.TotalMiles += trip.Miles
			driver.TotalHours += trip.EndTime.Sub(trip.StartTime).Hours()
			drivers[driver.Name] = driver
		}
	}

	for _, driver := range drivers {
		output := fmt.Sprintf("%s: %d miles", driver.Name, driver.RoundedTotalMiles())
		if driver.TotalMiles > 0 {
			output += fmt.Sprintf(" @ %d mph", driver.AverageMPH())
		}
		fmt.Println(output)
	}
}

func NewTrip(record []string) (Trip, error) {
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

	return Trip{
		StartTime: startTime,
		EndTime:   endTime,
		Miles:     miles,
	}, nil
}

func (d Driver) RoundedTotalMiles() int {
	return int(math.Round(float64(d.TotalMiles)))
}

func (d Driver) AverageMPH() int {
	if d.TotalMiles == 0 {
		return 0
	}

	return int(math.Round(float64(d.TotalMiles / d.TotalHours)))
}

type Driver struct {
	Name       string
	Trips      []Trip
	TotalMiles float64
	TotalHours float64
}

type Trip struct {
	StartTime time.Time
	EndTime   time.Time
	Miles     float64
}

type processedRecord struct {
	record []string
	err    error
}

func processCSV(rc io.Reader) chan processedRecord {
	ch := make(chan processedRecord)

	go func() {
		r := csv.NewReader(rc)
		r.Comma = ' '

		defer close(ch)
		for {
			rec, err := r.Read()

			if err != nil {

				//if error is EOF, we're done
				if err == io.EOF {
					break
				}

				err = func(err error) error {
					errParse, ok := err.(*csv.ParseError)
					if !ok {
						// unexpected error
						return err
					}

					// we are parsing two different data sets in one file,
					// so we will see this error for our two data sets
					// if the read record has a count we are expecting, we should process it
					if errParse.Err == csv.ErrFieldCount && (len(rec) == 5 || len(rec) == 2) {
						return nil
					}

					//this is truly an error that we are not expecting
					return err
				}(err)
			}

			ch <- processedRecord{
				record: rec,
				err:    err,
			}
		}

	}()

	return ch
}
