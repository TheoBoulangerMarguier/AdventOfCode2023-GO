package main

import (
	"AdventOfCode/Day23"
	"fmt"
)

func main() {

	results := [][2]int{
		{-1, -1}, //Day1.Day1(),
		{-1, -1}, //Day2.Day2(),
		{-1, -1}, //Day3.Day3(),
		{-1, -1}, //Day4.Day4(),
		{-1, -1}, //Day5.Day5(), //slow
		{-1, -1}, //Day6.Day6(),
		{-1, -1}, //Day7.Day7(),
		{-1, -1}, //Day8.Day8(),
		{-1, -1}, //Day9.Day9(),
		{-1, -1}, //Day10.Day10(),
		{-1, -1}, //Day11.Day11(),
		{-1, 0},  //Day12.Day12(), //only p1 found so far will come back later for p2
		{-1, -1}, //Day13.Day13(),
		{-1, -1}, //Day14.Day14(),
		{-1, -1}, //Day15.Day15(),
		{-1, -1}, //Day16.Day16(), //slow
		{-1, -1}, //Day17.Day17(), //slow
		{-1, -1}, //Day18.Day18(),
		{-1, 0},  //Day19.Day19(), //only p1 found so far will come back later for p2
		{-1, 0},  //Day20.Day20(), //only p1 found so far will come back later for p2
		{-1, -1}, //Day21.Day21(),
		{-1, -1}, //Day22.Day22(), //slightly slow but very fine
		Day23.Day23(),
	}

	testResults(results)
}

func testResults(result [][2]int) {
	expectedResults := [][2]int{
		{54916, 54728},
		{2176, 63700},
		{546312, 87449461},
		{26426, 6227972},
		{322500873, 108956227},
		{2374848, 39132886},
		{253866470, 254494947},
		{17873, 15746133679061},
		{1987402313, 900},
		{6886, 371},
		{10289334, 649862989626},
		{7047, -1},
		{35521, 34795},
		{106517, 79723},
		{502139, 284132},
		{7927, 8246},
		{785, 922},
		{74074, 112074045986829},
		{432788, -1},
		{788848550, -1},
		{3600, 599763113936220},
		{441, 80778},
		{-1, -1},
	}

	for i := 0; i < len(expectedResults); i++ {
		if expectedResults[i][0] == result[i][0] {
			fmt.Println("\033[32mDay: ", i+1, ", Part: 1, PASS result :", result[i][0], " matching expected result ", expectedResults[i][0], "\033[0m")
		} else if result[i][0] == -1 {
			fmt.Println("Day: ", i+1, ", Part: 1, SKIPED expected result ", expectedResults[i][0])
		} else {
			fmt.Println("\033[31mDay: ", i+1, ", Part: 1, FAIL result :", result[i][0], " not matching expected result ", expectedResults[i][0], "\033[0m")
		}

		if expectedResults[i][1] == result[i][1] {
			fmt.Println("\033[32mDay: ", i+1, ", Part: 2, PASS result :", result[i][1], " matching expected result ", expectedResults[i][1], "\033[0m")
		} else if result[i][1] == -1 {
			fmt.Println("Day: ", i+1, ", Part: 2, SKIPED expected result ", expectedResults[i][1])
		} else {
			fmt.Println("\033[31mDay: ", i+1, ", Part: 2, FAIL result :", result[i][1], " not matching expected result ", expectedResults[i][1], "\033[0m")
		}
	}
}
