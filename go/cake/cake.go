package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose  bool
	Cakes    int
	BakeTime time.Duration
	// Standard deviation (spread from average value) of baking time
	BakeStdDev time.Duration
	// Buffer slot between baking and icing
	BakeBuf int
	// Could be higher depending on the benchmarking context
	NumIcers       int
	IceTime        time.Duration
	IceStdDev      time.Duration
	IceBuf         int
	InscribeTime   time.Duration
	InscribeStdDev time.Duration
}

type cake int

// Run the simulation in 'runs' time
func (s *Shop) Work(runs int) {
	for r := 0; r < runs; r++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// Block the calling goroutine for a period of time
// that is normally distributed around mean
// with a standard deviation of stddev.
// We simulate work duration with natural variance,
// such as network latency, processing jitter, or randomized backoff.
func work(mean, stddev time.Duration) {
	delay := mean + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
