package main

import (
	"fmt"
	"math/rand"
	"time"
)

var secretNum int

func SecretNum(n int) {
	r := rand.New(rand.NewSource(time.Now().Unix())) // new way of using the rand.New ong go.20
	secretNum = r.Intn(n) + 1
	// we are adding 1 to the secretNum here because our rand.Intn find a random between 0 to n-1
}

func guess(n int) int {
	if n == secretNum {
		return 0
	} else if n < secretNum {
		return 1
	} else {
		return -1
	}
}

func guessNumber(n int) int {
	l, r := 1, n

	for l <= r {
		m := l + (r-l)/2
		res := guess(m)
		if res == 0 {
			return m
		} else if res < 0 {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func main() {
	n := 10
	SecretNum(n) // Guessing from nos 1- 10
	dGuessedNum := guessNumber(n)
	fmt.Println(dGuessedNum)
}
