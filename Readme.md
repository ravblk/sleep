## Sleep functions

Package sleep implements with various functions that you can to do when repeating.

Sleep functions in ms

```
sleep.MS(2)
```

Sleep with jitter in ms (jitter 0-1 ms)

```
sleep.MSJitter(1)
```

Random sleep 0-max ms

```
sleep.MSRandom(max)
```

Exponention sleep 2->4->8... ms

```
   power := 0
	for i := 0; i < 3; i++ {
 		power = MSExponent(power)
	}
```

Exponention sleep 2->4->8... ms with jitter(0-1ms)

```
   power = 0
	for i := 0; i < 3; i++ {
  		power = MSExponentJitter(power)
	}
```