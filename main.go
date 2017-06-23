package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func calcMonteCarloForChordLargerThanInscribedEquilateralTriangle(n int, ch chan<- int) {
	var r float64 = 1
	a := 0
	for i := 0; i < n; i++ {
		rand := rand.Float64()
		c := 2 * r * math.Sin(rand * math.Pi / 2)
		if c > r * math.Sqrt(3) {
			a += 1
		}
	}
	ch <- a
}



func main() {
	st := time.Now().UnixNano()
	a := 0
	var c chan int = make(chan int)

	parallel := 16
	n := 10000000

	for i := 0; i < parallel; i++ {
		go calcMonteCarloForChordLargerThanInscribedEquilateralTriangle(n, c)
	}

	for i := 0; i < parallel; i++ {
		a += <-c
	}

	fmt.Printf("a: %d\n", a)
	fmt.Printf("probability: %f\n", float64(a) / float64(n*parallel))
	fmt.Printf("Processed in %d seconds\n", (time.Now().UnixNano()-st)/int64(time.Second))
}
