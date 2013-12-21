package main

import (
	"fmt"
	"math/rand"
)

func gen(n int, seed int64) []float64 {
	rnd := rand.New(rand.NewSource(seed))
	res := make([]float64, n)
	for i := range res {
		res[i] = rnd.Float64()
	}
	return res
}

func count(p []float64) int {
	cnt := 0
	max := 1.01
	for _, v := range p {
		if v < max {
			cnt++
			max = v
		}
	}
	return cnt
}

func e(n, tries int) float64 {
	var cnt float64
	for i := 0; i < tries; i++ {
		cnt += float64(count(gen(n, int64(i))))
	}
	return cnt / float64(tries)
}

func main() {
	var a [20]float64
	for i := range a {
		a[i] = e(i, 100000)
		if i == 0 {
			continue
		}
		fmt.Printf("e(%d) = %f, e(%d) - e(%d) = %f\n", i, a[i], i, i-1, a[i]-a[i-1])
	}
}
