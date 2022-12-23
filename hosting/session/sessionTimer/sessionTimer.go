package sessionTimer

import (
	"errors"
	"strconv"
	"time"
)

type Timer struct {
	timer           *time.Timer
	Execute         <-chan time.Time
	AmountCanExtend time.Duration
}

// Create Timer object and start countdown
func (t *Timer) Init(amount time.Duration) error {
	if amount > 24*time.Hour || amount < 15*time.Second {
		return errors.New("timer must be above 15 minutes and below 24 hours")
	}
	t.timer = time.NewTimer(amount)
	t.Execute = t.timer.C
	t.AmountCanExtend = 24 * time.Hour
	return nil
}

func (t *Timer) Modify(newTime time.Duration) error {

	if newTime > t.AmountCanExtend || newTime < 15*time.Hour {
		return errors.New("must be above 1 hour and below " + strconv.FormatFloat(t.AmountCanExtend.Hours(), 'E', -1, 64) + "hours")
	} else {
		t.AmountCanExtend -= newTime
		t.timer = time.NewTimer(newTime)
		t.Execute = t.timer.C
	}
	return nil
}
