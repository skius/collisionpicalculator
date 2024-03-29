package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	var digits int
	if len(os.Args) > 1 {
		digits, _ = strconv.Atoi(os.Args[1])
	} else {
		digits = 10
	}

	b1 := Block{
		m: 1,
		v: 0,
	}
	b2 := Block{
		m: math.Pow(100, float64(digits-1)),
		v: -1,
	}

	var count int64 = 0
	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for range ticker.C {
			fmt.Println(count)
		}
	}()
	before := time.Now()
	bounce := true
	for {
		if b1.v >= 0 && b2.v > b1.v {
			break
		}
		if bounce {
			Bounce(&b1, &b2)
			//fmt.Printf("Ratio: %.10f\n", b2.v/b1.v)
		} else {
			b1.Reverse()
		}
		bounce = !bounce
		count++
	}

	after := time.Now()
	passed := after.Sub(before).Nanoseconds() / 1000 / 1000

	fmt.Println("Calculated", digits, "digits of Pi:", float64(count)/math.Pow(10, float64(digits-1)), "in", passed, "ms.")
}
