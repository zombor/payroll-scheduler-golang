package main

import "testing"
import "time"

func TestNextPayDateWithWeekFrequency(t *testing.T) {
  in := time.Date(2014, time.February, 13, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 20, 0, 0, 0, 0, time.UTC)

  if x := NextPayDate(in, 1); x != out {
    t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSaturday(t *testing.T) {
  in := time.Date(2014, time.February, 8, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x := NextPayDate(in, 2); x != out {
    t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
  }
}

func TestFallsBackToFridayIfPaydayIsOnSunday(t *testing.T) {
  in := time.Date(2014, time.February, 9, 0, 0, 0, 0, time.UTC)
  out := time.Date(2014, time.February, 21, 0, 0, 0, 0, time.UTC)

  if x := NextPayDate(in, 2); x != out {
    t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
  }
}
