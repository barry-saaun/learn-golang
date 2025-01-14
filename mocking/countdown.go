package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

// SpySleeper for testing

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpySleeper struct {
	Calls int
}

// this records how many time (sleep per second) the incrementation has been called
// to get rid of the overhead in waiting "n" of seconds in testing
func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOpetations struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownOpetations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOpetations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Millisecond, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
