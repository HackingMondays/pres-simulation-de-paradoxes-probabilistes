package main

import (
	"fmt"
	"math"
	"math/rand"
)

const MIN_SIZE_GUESS = 1
const MAX_SIZE_GUESS = 100
const MAX_PROBA_ITER = 10000

const TARGET_PCT = 50

func main() {
	fmt.Println("To have 50% chance of same birthday, choose group of size:", findSize())
}

func findSize() int {
	a := MIN_SIZE_GUESS
	b := MAX_SIZE_GUESS

	// apply bisection method
	for (b - a) > 1 {
		m := (a + b) / 2
		fa, fm := calculateRange(a, m)
		if fa*fm > 0 {
			a = m
		} else {
			b = m
		}
	}
	// now decide between interval boundaries [a,b]
	fa := probaForSize(a) - TARGET_PCT
	fb := probaForSize(b) - TARGET_PCT

	return smallest(a, b, fa, fb)
}

// return boundary value with smallest delta to target percentage
func smallest(a, b, fa, fb int) int {
	if fa*fa < fb*fb {
		return a
	} else {
		return b
	}
}

// interesting way of making a function concurrent
// https://talks.golang.org/2012/concurrency.slide#25
func calculateRange(a int, b int) (int, int) {
	c := make(chan int)
	go func() { c <- probaForSize(a) }()
	go func() { c <- probaForSize(b) }()
	return (<-c - TARGET_PCT), (<-c - TARGET_PCT)
}

// calculate probability for a group of specified size
func probaForSize(size int) int {
	var count int
	for i := 1; i < MAX_PROBA_ITER; i++ {
		if randomGroupHasDuplicates(size) {
			count++
		}
	}
	return count * 100 / MAX_PROBA_ITER
}

// return true if two people have the same birthday
// in a random group of specified size
func randomGroupHasDuplicates(size int) bool {
	var birthdayOn [365]bool
	for i := 0; i < size; i++ {
		date := int(math.Floor(rand.Float64() * 365))
		if birthdayOn[date] {
			return true
		}
		birthdayOn[date] = true
	}
	return false
}
