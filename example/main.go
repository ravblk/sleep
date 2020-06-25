package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------sleep time--------------")
	start := time.Now()
	sleep := 2
	MS(sleep)
	fmt.Println("sleep time", sleep, "ms")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------sleep jitter time--------------")
	start = time.Now()
	sleep = 5
	MSJitter(sleep)
	fmt.Println("sleep time without jitter", sleep, "ms")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------random sleep time--------------")
	start = time.Now()
	max := 20
	MSRandom(max)
	fmt.Println("sleep max time", max, "ms")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------sleep exponent time--------------")
	power := 0
	for i := 0; i < 3; i++ {
		start = time.Now()
		power = MSExponent(power)
		fmt.Println("measured time", time.Since(start))
		fmt.Println("power", power)
	}

	fmt.Println("--------------sleep exponent and jitter time--------------")
	power = 0
	for i := 0; i < 3; i++ {
		start = time.Now()
		power = MSExponentJitter(power)
		fmt.Println("measured time", time.Since(start))
		fmt.Println("power without jitter", power)
	}

	fmt.Println("--------------sleep time--------------")
	start = time.Now()
	sleep = 1
	Second(sleep)
	fmt.Println("sleep time", sleep, "second")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------sleep jitter time--------------")
	start = time.Now()
	sleep = 1
	SecondJitter(sleep)
	fmt.Println("sleep time without jitter", sleep, "second")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------random sleep time--------------")
	start = time.Now()
	max = 4
	SecondRandom(max)
	fmt.Println("sleep max time", max, "second")
	fmt.Println("measured time", time.Since(start))

	fmt.Println("--------------sleep exponent time--------------")
	power = 0
	for i := 0; i < 2; i++ {
		start = time.Now()
		power = SecondExponent(power)
		fmt.Println("measured time", time.Since(start))
		fmt.Println("power", power)
	}

	fmt.Println("--------------sleep exponent and jitter time--------------")
	power = 0
	for i := 0; i < 2; i++ {
		start = time.Now()
		power = SecondExponentJitter(power)
		fmt.Println("measured time", time.Since(start))
		fmt.Println("power without jitter", power)
	}
}
