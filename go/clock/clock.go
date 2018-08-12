package clock

import "fmt"

const (
    HOURS_IN_DAY = 24
    MINUTES_IN_HOUR = 60
    MINUTES_IN_DAY = MINUTES_IN_HOUR * HOURS_IN_DAY
)

type Clock struct {
    minutes int
}

func(c *Clock) Normalize() {
    // go modulus 
    // -a%b == -(a%b)
    c.minutes %= MINUTES_IN_DAY
    if c.minutes < 0 {
        c.minutes += MINUTES_IN_DAY
    }
}

func New(hour, minute int) Clock {
    minutes := minute + (MINUTES_IN_HOUR * hour)
    c := Clock{ minutes }
    c.Normalize()
    return c
}

func (c Clock) String() string {
    h := c.minutes / MINUTES_IN_HOUR
    m := c.minutes % MINUTES_IN_HOUR

    timeString := fmt.Sprintf("%02d:%02d", h, m)
    return timeString
}

func (c Clock) Add(minutes int) Clock {
    c.minutes += minutes
    c.Normalize()
    return c
}

func (c Clock) Subtract(minutes int) Clock {
    c.minutes -= minutes
    c.Normalize()
    return c
}
