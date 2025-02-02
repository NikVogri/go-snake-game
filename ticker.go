package main

import (
	"time"
)

type Ticker struct {
	Speed time.Duration
	Tick  bool
}

func createTicker() *Ticker {
	return &Ticker{
		Speed: 1,
		Tick:  false,
	}
}

func (ticker *Ticker) start(fn func(t *Ticker)) {
	ticker.Tick = true

	for ticker.Tick {
		fn(ticker)
		time.Sleep(ticker.Speed)
	}
}

func (ticker *Ticker) stop() {
	ticker.Tick = false
}

func (ticker *Ticker) setSpeed(speed time.Duration) {
	ticker.Speed = speed
}

func (ticker *Ticker) increaseSpeedByFactor(factor float64) {
	ticker.Speed = ticker.Speed - time.Duration((float64(ticker.Speed) * factor))
}
