package sessionTimer

import (
	"fmt"
	"time"
)

type Timer struct {
	time *time.Timer
}

func init() {
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

}

func (t *Timer) Create(amount time.Duration) *Timer {
	temp := &Timer{time: time.NewTimer(amount)}
	return temp
}
