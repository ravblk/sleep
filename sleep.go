package sleep

import (
	"math/rand"
	"time"
)

const (
	jitterMkS = 1000
	jitterMS  = 1000
)

//MS sleep  d ms
func MS(d int) {
	time.Sleep(time.Duration(d) * time.Millisecond)
}

//MSJitter sleep d + jitter 0-1ms
func MSJitter(d int) {
	dj := rand.Int() % jitterMkS
	time.Sleep(time.Duration(d)*time.Millisecond + time.Duration(dj)*time.Microsecond)
}

//MSRandom sleep random number  0 <= n < max ms
func MSRandom(max int) {
	d := rand.Int() % max
	MS(d)
}

//MSExponent sleep 2^power ms return d squared
func MSExponent(power int) int {
	MS(2 << power)
	power++

	return power
}

//MSExponentJitter sllep 2^power + jitter 0-1ms and return d squared
func MSExponentJitter(power int) int {
	MSJitter(2 << power)
	power++

	return power
}

//Second sleep  d second
func Second(d int) {
	time.Sleep(time.Duration(d) * time.Second)
}

//SecondJitter sleep d second + jitter 0-1second
func SecondJitter(d int) {
	dj := rand.Int() % jitterMS
	time.Sleep(time.Duration(d)*time.Second + time.Duration(dj)*time.Millisecond)
}

//SecondRandom sleep random number  0 <= n < max second
func SecondRandom(max int) {
	d := rand.Int() % max
	Second(d)
}

//SecondExponent sleep 2^power second return d squared
func SecondExponent(power int) int {
	Second(2 << power)
	power++

	return power
}

//SecondExponentJitter sleep 2^power second + jitter 0-1s and return d squared
func SecondExponentJitter(power int) int {
	SecondJitter(2 << power)
	power++

	return power
}
