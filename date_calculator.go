package main

import "time"
import "errors"

func NextPayDate(start_date time.Time, num_weeks int32) (time.Time, error) {
  original_year := start_date.Year()

  date := start_date.Add(time.Duration(num_weeks*24*7)*time.Hour)

  for (Weekend(date.Weekday()) == true) {
    date = date.Add(-time.Duration(24*time.Hour))
  }

  if date.Year() != original_year {
    return date, errors.New("Cannot iterate past current year")
  }

  return date, nil
}

func Weekend(weekday time.Weekday) bool {
  return (weekday == 0 || weekday == 6)
}
