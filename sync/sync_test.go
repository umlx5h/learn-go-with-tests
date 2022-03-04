package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leave it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("incrementing the counter concurrently", func(t *testing.T) {
		counter := NewCounter()
		countNum := 100000

		var wg sync.WaitGroup
		wg.Add(countNum)
		for i := 0; i < countNum; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}

		wg.Wait()

		assert.Equal(t, countNum, counter.Value())
	})
}
