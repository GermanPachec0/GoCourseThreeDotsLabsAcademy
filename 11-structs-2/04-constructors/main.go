package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func NewDateRange(start, end time.Time) (*DateRange, error) {
	if start.IsZero() {
		return nil, errors.New("Las fechas no pueden ser null.")
	}
	if end.IsZero() {
		return nil, errors.New("Las fechas no pueden ser null.")
	}
	if end.Before(start) {
		return nil, errors.New("La fecha de fin no puede ser antes de la de comienzo.")
	}
	return &DateRange{
		Start: start,
		End:   end,
	}, nil
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	start := time.Date(2022, time.January, 1, 12, 30, 0, 0, time.UTC)
	end := time.Date(2022, time.January, 2, 15, 30, 0, 0, time.UTC)
	lifetime, err := NewDateRange(start, end)

	if err != nil {
		fmt.Print("Error Creating Date")
	}

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
