package payroll_scheduler

import "time"

func Weekend(weekday time.Weekday) bool {
  return (weekday == 0 || weekday == 6)
}
