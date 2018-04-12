package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

type processedRecord struct {
	record []string
	err    error
}

//process reads the reader and populates the drivers map
func process(reader io.Reader, drivers map[string]Driver) error {
	if reader == nil {
		return errors.New("reader is nil")
	}
	for rec := range read(reader, ' ') {
		err := processRecord(drivers, rec.record)
		if err != nil {
			//decision was made to just print out to screen if there is an error in the record
			//this could be return in a chan back to the caller so they could handle the errors
			fmt.Printf("Processing line : %v : had an error : %v \n", rec.record, err)
			continue
		}
	}

	return nil
}

//read parses each line using the delimiter and writes it to a channel
func read(rc io.Reader, delimiter rune) chan processedRecord {
	ch := make(chan processedRecord)

	go func() {
		r := csv.NewReader(rc)
		r.Comma = delimiter

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
					if errParse.Err == csv.ErrFieldCount {
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

func processRecord(drivers map[string]Driver, record []string) error {

	//if there is no data in the record, return an error

	//An assumption on the data we have here is that every command will at least have two fields
	if len(record) < 2 {
		return errors.New("Data invalid in record")
	}

	switch strings.TrimSpace(record[0]) {
	case "Driver":
		//An assumption on the data here is that each data point will have driver as the second field
		//No need to account for trips that do not have registered drivers, so create a new one
		//if the command is a Trip or Driver
		name := strings.TrimSpace(record[1])
		driver, ok := drivers[name]
		if !ok {
			driver = Driver{
				Name:  name,
				Trips: make([]Trip, 0),
			}
		}
		drivers[driver.Name] = driver
	case "Trip":
		driver, ok := drivers[strings.TrimSpace(record[1])]
		if !ok {
			return fmt.Errorf("No driver registered for this trip : %s", strings.TrimSpace(record[1]))
		}

		trip, err := NewTrip(record)
		if err != nil {
			return err
		}

		driver.Trips = append(driver.Trips, trip)
		driver.TotalMiles += trip.Miles
		driver.TotalHours += trip.TotalHours
		drivers[driver.Name] = driver
	default:
		return fmt.Errorf("Unexpected command : %s", strings.TrimSpace(record[0]))
	}

	return nil
}
