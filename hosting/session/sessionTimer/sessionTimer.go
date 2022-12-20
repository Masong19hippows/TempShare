package sessionTimer

import (
	"errors"
	"strconv"
	"time"
)

type Timer struct {
	timer           *time.Timer
	execute         <-chan time.Time
	AmountCanExtend time.Duration
}

func (t *Timer) initTimer() {
	t.execute = t.timer.C
}

// Create Timer object and start countdown
func Create(amount time.Duration) (*Timer, error) {
	if amount > 24*time.Hour || amount < 1*time.Hour {
		return nil, errors.New("must be above 1 hour and below 24 hours")
	}
	temp := &Timer{timer: time.NewTimer(amount)}
	temp.initTimer()
	temp.AmountCanExtend = 24 * time.Hour

	return temp, nil
}

func (t *Timer) Modify(newTime time.Duration) error {

	if newTime > t.AmountCanExtend || newTime < 15*time.Hour {
		return errors.New("must be above 1 hour and below " + strconv.FormatFloat(t.AmountCanExtend.Hours(), 'E', -1, 64) + "hours")
	} else {
		t.AmountCanExtend -= newTime
		t.timer = time.NewTimer(newTime)
		t.initTimer()
	}
	return nil
}
