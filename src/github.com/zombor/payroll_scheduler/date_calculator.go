package payroll_scheduler

import "time"
import "errors"

func NextWeeklyPayDate(start_date time.Time, num_weeks int32) (time.Time, error) {
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

func NextMonthlyPayDate(start_date time.Time) (time.Time, error) {
  month := start_date.Month()
  next_month := month+1

  if (next_month > 12) {
    return start_date, errors.New("Cannot iterate past current year")
  }

  next_date := time.Date(start_date.Year(), next_month, start_date.Day(), 0, 0, 0, 0, time.UTC)

  for (next_date.Month() != next_month) {
    next_date = next_date.Add(-time.Duration(24*time.Hour))
  }

  for (Weekend(next_date.Weekday()) == true) {
    next_date = next_date.Add(-time.Duration(24*time.Hour))
  }

  return next_date, nil
}
