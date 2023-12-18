package main

import (
	"AdventOfCode/Day16"
	"AdventOfCode/Day5"

	"testing"
)

func BenchmarkDay5(b *testing.B) {
	Day5.Day5()
}

func BenchmarkDay16(b *testing.B) {
	Day16.Day16()
}
