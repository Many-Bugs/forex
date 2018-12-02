package times

import (
	"fmt"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
)

type Schedular struct {
	clauses map[string]interface{}
	errs    map[string][]error
}

func New() *Schedular {
	var schedular Schedular
	schedular.clauses["retry"] = 10
	return &schedular
}

func ParseAny(value string) {
	t, err := dateparse.ParseAny(value)
	fmt.Println(t, err)
}

func (s *Schedular) timeTypeConditioning(param string, value interface{}) *Schedular {
	now := time.Now()
	switch value.(type) {
	case string:
		t, err := dateparse.ParseAny(value.(string))
		if err != nil {
			s.clauses[param] = err
		}
		s.clauses[param] = t
	case time.Time:
		s.clauses[param] = value
	case [4]int:
		s.clauses[param] = time.Date(now.Year(), now.Month(), now.Day(), value.([4]int)[0], value.([4]int)[1], value.([4]int)[2], value.([4]int)[3], time.Local)
	default:
		s.clauses[param] = now
	}
	return s
}

func (s *Schedular) intervalTypeConditioning(value interface{}) *Schedular {
	var interval time.Duration
	switch value.(type) {
	case string:
		t, err := dateparse.ParseAny(value.(string))
		if err != nil {
			period, err := strconv.Atoi(value.(string))
			if err != nil {
				s.errs["interval"] = append(s.errs["interval"], err)
				return s
			}
			s.clauses["interval"] = time.Duration(period)
			return s
		}
		interval = time.Since(t)
	case int:
		interval = time.Duration(value.(int))
	case [2]time.Time:
		interval = (value.([2]time.Time)[1]).Sub(value.([2]time.Time)[0])
	case [2]string:
		t0, err0 := dateparse.ParseAny((value.([2]string)[0]))
		if err0 != nil {
			s.errs["interval"] = append(s.errs["interval"], err0)
			return s
		}
		t1, err1 := dateparse.ParseAny((value.([2]string)[1]))
		if err1 != nil {
			s.errs["interval"] = append(s.errs["interval"], err1)
			return s
		}
		interval = t1.Sub(t0)
	case time.Duration:
		interval = value.(time.Duration)
	default:
		s.errs["interval"] = append(s.errs["interval"], fmt.Errorf("invalid interval condition: %v", value))
		return s
	}
	s.clauses["interval"] = interval
	return s
}

func (s *Schedular) timetableConditioning(value ...interface{}) *Schedular {

	return s
}

func (s *Schedular) Start(value interface{}) *Schedular {
	s.timeTypeConditioning("start", value)
	return s
}

func (s *Schedular) End(value interface{}) *Schedular {
	s.timeTypeConditioning("end", value)
	return s
}

func (s *Schedular) Retry(count int) *Schedular {
	s.clauses["retry"] = count
	return s
}

func (s *Schedular) Interval(value interface{}) *Schedular {
	s.intervalTypeConditioning(value)
	return s
}

func (s *Schedular) Function(f func(...interface{}) error) *Schedular {
	s.clauses["function"] = f
	return s
}

func (s *Schedular) Params(params ...interface{}) *Schedular {
	s.clauses["params"] = params
	return s
}

func (s *Schedular) Timetable(tables ...interface{}) *Schedular {
	s.timetableConditioning(tables)
	return s
}

func (s *Schedular) Excute() *Schedular {
	return s
}

// count use for retry count, interval use for scheduling e.g. 4 as 4 hours
func Routine(start []int, count int, interval time.Duration, f func() error) {
	if len(start) > 4 || len(start) == 0 {
		start = []int{0, 0, 0, 0}
	}

	if err := f(); err != nil {
		retry(1*time.Minute, count, f)
	}

	ticker := updateTicker(start, interval)
	for {
		<-ticker.C
		if err := f(); err != nil {
			retry(1*time.Minute, count, f)
		}
		ticker = updateTicker(start, interval)
	}
}

// count use for retry count, interval use for scheduling, interval has to be factor of 24 and start - end e.g. 4 as 4 hours
func RoutineOnSchedule(start, end []int, count int, interval time.Duration, f func() error) {

	if len(start) > 4 || len(start) == 0 {
		start = []int{0, 0, 0, 0}
	}

	if err := f(); err != nil {
		retry(1*time.Minute, count, f)
	}

	ticker := updateTicker(start, interval)
	for {
		<-ticker.C
		now := time.Now().Local()
		s := time.Date(now.Year(), now.Month(), now.Day(), start[0], start[1], start[2], start[3], time.Local)
		e := time.Date(now.Year(), now.Month(), now.Day(), end[0], end[1], end[2], end[3], time.Local)
		if now.After(s) && now.Before(e) {
			if err := f(); err != nil {
				retry(1*time.Minute, count, f)
			}
		}
		ticker = updateTicker(start, interval)
	}
}

func updateTicker(start []int, interval time.Duration) *time.Ticker {
	now := time.Now().Local()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), start[0], start[1], start[2], start[3], time.Local)
	for {
		if !nextTick.After(now) {
			nextTick = nextTick.Add(interval)
		} else {
			break
		}
	}

	diff := nextTick.Sub(now)
	return time.NewTicker(diff)
}

func retry(duration time.Duration, count int, f func() error) {
	for range time.Tick(duration) {
		f()
		count--
		if count <= 0 {
			break
		}
	}
}
