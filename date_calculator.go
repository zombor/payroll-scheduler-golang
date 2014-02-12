package main

import "time"

func NextPayDate(start_date time.Time, num_weeks int32) time.Time {
  date := start_date.Add(time.Duration(num_weeks*24*7)*time.Hour)

  for (Weekend(date.Weekday()) == true) {
    date = date.Add(-time.Duration(24*time.Hour))
  }

  return date
}

func Weekend(weekday time.Weekday) bool {
  return (weekday == 0 || weekday == 6)
}
