package main

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name        string
		input       io.Reader
		expected    map[string]Driver
		errExpected error
	}{
		{
			name: "Success",
			input: strings.NewReader(`
					Driver Dan
					Trip Dan 07:15 07:45 10.0`),
			expected: map[string]Driver{
				"Dan": Driver{
					Name:       "Dan",
					TotalMiles: 10.0,
					TotalHours: 0.5,
					Trips: []Trip{
						Trip{
							StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
							EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
							TotalHours: 0.5,
							Miles:      10.0,
						},
					},
				},
			},
			errExpected: nil,
		},
		{
			name:        "Invalid command",
			input:       strings.NewReader("Invalid Dan"),
			expected:    map[string]Driver{},
			errExpected: nil,
		},
		{
			name:        "Invalid input",
			input:       strings.NewReader("Invalid"),
			expected:    map[string]Driver{},
			errExpected: nil,
		},
		{
			name:        "Invalid reader",
			input:       nil,
			expected:    map[string]Driver{},
			errExpected: errors.New("reader is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := map[string]Driver{}
			err := process(tt.input, actual)
			if err == tt.errExpected && reflect.TypeOf(err) != reflect.TypeOf(tt.errExpected) {
				t.Errorf("failed at %s : expected error %v : actual error : %v", tt.name, tt.errExpected, err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("failed at %s : expected %+v : actual : %+v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestProcessRecord(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		errExpected error
	}{
		{
			name:        "Success",
			input:       []string{"Driver", "Dan"},
			errExpected: nil,
		},
		{
			name:        "Driver not registered",
			input:       []string{"Trip", "Dan", "07:15", "07:45", "10.0"},
			errExpected: errors.New("Driver not registered"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := processRecord(map[string]Driver{}, tt.input)

			if err != tt.errExpected && reflect.TypeOf(err) != reflect.TypeOf(tt.errExpected) {
				t.Errorf("failed at %s : expected error %v : actual error : %v", tt.name, tt.errExpected, err)
			}
		})
	}
}
