package main

import "time"

func NextPayDate(start_date time.Time, num_weeks int32) time.Time {
  date := start_date.Add(time.Duration(num_weeks*24*7)*time.Hour)
  if (date.Weekday().String() == "Saturday") {
    return date.Add(-time.Duration(24*time.Hour))
  } else if (date.Weekday().String() == "Sunday") {
    return date.Add(-time.Duration(24*2*time.Hour))
  } else {
    return date
  }
}
