package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSortDrivers(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]Driver
		expected []Driver
	}{
		{
			name: "Success",
			input: map[string]Driver{
				"Dan": {
					Name:       "Dan",
					TotalMiles: 10.0,
					TotalHours: 0.5,
					Trips: []Trip{
						{
							StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
							EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
							TotalHours: 0.5,
							Miles:      10.0,
						},
					},
				},
				"Alex": {
					Name:       "Alex",
					TotalMiles: 5.0,
					TotalHours: 0.5,
					Trips: []Trip{
						{
							StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
							EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
							TotalHours: 0.5,
							Miles:      5.0,
						},
					},
				},
			},
			expected: []Driver{
				{
					Name:       "Dan",
					TotalMiles: 10.0,
					TotalHours: 0.5,
					Trips: []Trip{
						{
							StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
							EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
							TotalHours: 0.5,
							Miles:      10.0,
						},
					},
				},
				{
					Name:       "Alex",
					TotalMiles: 5.0,
					TotalHours: 0.5,
					Trips: []Trip{
						{
							StartTime:  time.Date(0, 1, 1, 7, 15, 0, 0, time.UTC),
							EndTime:    time.Date(0, 1, 1, 7, 45, 0, 0, time.UTC),
							TotalHours: 0.5,
							Miles:      5.0,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sortDrivers(tt.input)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("failed at %s : expected %+v : actual : %+v", tt.name, tt.expected, actual)
			}
		})
	}
}
