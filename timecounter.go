package main

import (
	"fmt"
	"time"
	"timecounter/timescale"
)

var runningTime time.Time
var timeScale timescale.TimeScale
var timeSpent float64

// TimeCounter monitors the time elapsed between Start() and End()
func TimeCounter(ts timescale.TimeScale) {
	timeScale = ts
}

// Start time counting
func Start() {
	runningTime = time.Now()
}

// End finishes time counting
func End() {
	switch timeScale {
	case timescale.HOURS:
		timeSpent = time.Since(runningTime).Hours()
	case timescale.MINUTES:
		timeSpent = time.Since(runningTime).Minutes()
	case timescale.SECONDS:
		timeSpent = time.Since(runningTime).Seconds()
	case timescale.NANOSECONDS:
		t := float64(time.Since(runningTime).Nanoseconds())
		timeSpent = t
	case timescale.MILLISECONDS:
		fallthrough
	default:
		t := float64(time.Since(runningTime).Milliseconds())
		timeSpent = t
	}
}

// PrintTime print the time elapsed between Start() and End()
func PrintTime(s ...string) {
	if len(s) > 0 {
		fmt.Println(s)
		return
	}

	str := fmt.Sprintf("This code execution took -> %v %v", timeSpent, timeScale)
	fmt.Println(str)
}