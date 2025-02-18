package internal

import "time"

func NewRateLimiter(rps, rpm int) *time.Ticker {
	if rps == 0 && rpm == 0 {
		return nil
	}
	var interval time.Duration
	if rps > 0 {
		interval = time.Second / time.Duration(rps)
	} else if rpm > 0 {
		interval = time.Minute / time.Duration(rpm)
	}
	return time.NewTicker(interval)
}
