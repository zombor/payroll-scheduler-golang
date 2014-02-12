package main

import "testing"
import "time"

func TestNextPayDateWithWeekFrequency(t *testing.T) {
  in := time.Date(2014, time.February, 13, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 20, 0, 0, 0, 0, time.UTC)

  if x, _ := NextPayDate(in, 1); x != out {
    t.Errorf("NextPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSaturday(t *testing.T) {
  in := time.Date(2014, time.February, 8, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x, _ := NextPayDate(in, 2); x != out {
    t.Errorf("NextPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSunday(t *testing.T) {
  in := time.Date(2014, time.February, 9, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x, _ := NextPayDate(in, 2); x != out {
    t.Errorf("NextPayDate(%v) = %v, want %v", in, x, out)
  }
}

func TestErrorWhenRollsToNextYear(t *testing.T) {
  in := time.Date(2014, time.December, 26, 0, 0, 0, 0, time.UTC)

  if _, err := NextPayDate(in, 2); err == nil {
    t.Errorf("NextPayDate(%v) = %v, want not %v", in, err, nil)
  }
}
