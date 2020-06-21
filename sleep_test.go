package main

import (
	"testing"
	"time"

	"github.com/go-playground/assert"
)

func TestMS(t *testing.T) {
	start := time.Now()
	sleep := 2
	MS(sleep)
	assert.Equal(t, sleep, int(time.Since(start)/time.Millisecond))

	start = time.Now()
	sleep = 5
	MSJitter(sleep)
	res1 := assert.IsEqual(sleep, int(time.Since(start)/time.Millisecond))
	res2 := assert.IsEqual(sleep+1, int(time.Since(start)/time.Millisecond))

	assert.Equal(t, true, res1 || res2)

	start = time.Now()
	max := 20
	MSRandom(max)
	assert.Equal(t, true, max > int(time.Since(start)/time.Millisecond))

	power := 0
	for i := 0; i < 3; i++ {
		power = MSExponent(power)
		assert.Equal(t, i+1, power)
	}

	power = 0
	for i := 0; i < 3; i++ {
		power = MSExponentJitter(power)
		assert.Equal(t, i+1, power)
	}
}

func TestSec(t *testing.T) {
	start := time.Now()
	sleep := 2
	Second(sleep)
	assert.Equal(t, sleep, int(time.Since(start)/time.Second))

	start = time.Now()
	sleep = 1
	SecondJitter(sleep)
	assert.Equal(t, sleep, int(time.Since(start)/time.Second))

	start = time.Now()
	max := 5
	SecondRandom(max)
	assert.Equal(t, true, max > int(time.Since(start)/time.Second))

	power := 0
	for i := 0; i < 2; i++ {
		power = SecondExponent(power)
		assert.Equal(t, i+1, power)
	}

	power = 0
	for i := 0; i < 2; i++ {
		power = SecondExponentJitter(power)
		assert.Equal(t, i+1, power)
	}
}
