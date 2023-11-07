package main

import (
	"testing"
	"time"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	tolerance := 1000 * time.Millisecond // Погрешность в 100 миллисекунд

	tests := []struct {
		channels []<-chan interface{}
		want     time.Duration
	}{
		{[]<-chan interface{}{sig(2 * time.Hour), sig(5 * time.Minute), sig(1 * time.Second)}, 1 * time.Second},
		{[]<-chan interface{}{sig(1 * time.Hour), sig(1 * time.Minute)}, 1 * time.Minute},
	}

	for _, tt := range tests {
		start := time.Now()
		<-or(tt.channels...)
		got := time.Since(start)

		if got < tt.want-tolerance {
			t.Errorf("or() got = %v, want >= %v", got, tt.want-tolerance)
		}
	}
}
