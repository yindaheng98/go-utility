package SkipList

import (
	"math"
	"math/rand"
)

//This is the struct for the random generation of the number of level in SkipList.
//The detail of the algorithm: https://yindaheng98.github.io/%E6%95%B0%E5%AD%A6/SkipList.html
type RandLevel struct {
	logC    float64
	min     float64
	randSrc rand.Source
}

//Returns a pointer to a RandLevel.
//[index in level n]=[index in level n-1]/C, max output <= Level.
func NewRandomLevel(C, Level uint64, seed int64) *RandLevel {
	if C <= 1 {
		C = 2
	}
	c := float64(C)
	level := float64(Level)
	return &RandLevel{math.Log(c), 1.0 / math.Pow(c, level+1), rand.NewSource(seed)}
}

//Generate a random number.
func (rl *RandLevel) Rand() uint64 {
	X0 := float64(0)
	for X0 == 0 {
		X0 = rand.New(rl.randSrc).Float64()
	}
	X0 = X0*(1-rl.min) + rl.min
	X := math.Floor(-math.Log(X0) / rl.logC)
	return uint64(X)
}
