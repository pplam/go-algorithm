package filter

import (
	"sort"
)

func Filter(stop int) []int {
	if stop < 2 {
		return []int{}
	}

	result := []int{2}

	l := 2
	r := 2 * 2
	for stop > l {
		if stop < r {
			r = stop
		}

		ch := make(chan int)
		for x := l + 1; x <= r; x++ {
			go test(x, result, ch)
		}

		primes := []int{}
		counter := 0
		total := r - l
		for {
			if counter == total {
				break
			}

			prime := <-ch
			if prime > 0 {
				primes = append(primes, prime)
			}
			counter++
		}

		sort.Ints(primes)
		result = append(result, primes...)

		l = r
		r = l * l
	}

	return result
}

func test(x int, known []int, ch chan<- int) {
	for _, p := range known {
		if p * p > x {
			ch <- x
			return
		}
		if x % p == 0 {
			ch <- 0
			return
		}
	}
	ch <- x
}