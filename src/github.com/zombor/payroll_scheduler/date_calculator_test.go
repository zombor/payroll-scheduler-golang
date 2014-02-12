package payroll_scheduler

import "testing"
import "time"

func TestNextWeeklyPayDateWithWeekFrequency(t *testing.T) {
  in := time.Date(2014, time.February, 13, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 20, 0, 0, 0, 0, time.UTC)

  if x, _ := NextWeeklyPayDate(in, 1); x != out {
    t.Errorf("NextWeeklyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSaturday(t *testing.T) {
  in := time.Date(2014, time.February, 8, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x, _ := NextWeeklyPayDate(in, 2); x != out {
    t.Errorf("NextWeeklyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSunday(t *testing.T) {
  in := time.Date(2014, time.February, 9, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x, _ := NextWeeklyPayDate(in, 2); x != out {
    t.Errorf("NextWeeklyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestErrorWhenRollsToNextYear(t *testing.T) {
  in := time.Date(2014, time.December, 26, 0, 0, 0, 0, time.UTC)

  if _, err := NextWeeklyPayDate(in, 2); err == nil {
    t.Errorf("NextWeeklyPayDate(%v) = %v, want not %v", in, err, nil)
  }
}

func TestNextMonthlyPayDate(t *testing.T) {
  in := time.Date(2014, time.February, 13, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.March, 13, 0, 0, 0, 0, time.UTC)

  if x, _ := NextMonthlyPayDate(in); x != out {
    t.Errorf("NextMonthlyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestNextMonthlyPayDateForMissingDay(t *testing.T) {
  in := time.Date(2014, time.January, 30, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 28, 0, 0, 0, 0, time.UTC)

  if x, _ := NextMonthlyPayDate(in); x != out {
    t.Errorf("NextMonthlyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestNextMonthlyPayDateForMissingDayAndWeekend(t *testing.T) {
  in := time.Date(2014, time.July, 31, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.August, 29, 0, 0, 0, 0, time.UTC)

  if x, _ := NextMonthlyPayDate(in); x != out {
    t.Errorf("NextMonthlyPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestErrorWhenMonthlyRollsToNextYear(t *testing.T) {
  in := time.Date(2014, time.December, 26, 0, 0, 0, 0, time.UTC)

  if _, err := NextMonthlyPayDate(in); err == nil {
    t.Errorf("NextWeeklyPayDate(%v) = %v, want not %v", in, err, nil)
  }
}
