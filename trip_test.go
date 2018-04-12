package main

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNewTrip(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expected    Trip
		errExpected error
	}{
		{
			name:  "Success",
			input: []string{"Trip", "Dan", "07:15", "07:45", "10.0"},
			expected: Trip{
				StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
				EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
				TotalHours: 0.5,
				Miles:      10.0,
			},
			errExpected: nil,
		},
		{
			name:        "Invalid length",
			input:       []string{},
			expected:    Trip{},
			errExpected: errors.New("Invalid length"),
		},
		{
			name:        "Invalid start time",
			input:       []string{"Trip", "Dan", "0715", "07:45", "10.0"},
			expected:    Trip{},
			errExpected: errors.New("Invalid start time"),
		},
		{
			name:        "Invalid end time",
			input:       []string{"Trip", "Dan", "07:15", "0745", "10.0"},
			expected:    Trip{},
			errExpected: errors.New("Invalid end time"),
		},
		{
			name:        "Invalid miles",
			input:       []string{"Trip", "Dan", "07:15", "07:45", "one hundred"},
			expected:    Trip{},
			errExpected: errors.New("Invalid miles"),
		},
		{
			name:        "Trip was under 5 MPH",
			input:       []string{"Trip", "Dan", "07:15", "07:45", "2"},
			expected:    Trip{},
			errExpected: errors.New("under 5 MPH"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewTrip(tt.input)

			if err != tt.errExpected && reflect.TypeOf(err) != reflect.TypeOf(tt.errExpected) {
				t.Errorf("failed at %s : expected error %v : actual error : %v", tt.name, tt.errExpected, err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("failed at %s : expected %+v : actual : %+v", tt.name, tt.expected, actual)
			}
		})
	}
}
