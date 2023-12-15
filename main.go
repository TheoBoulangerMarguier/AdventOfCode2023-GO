package main

import (
	"AdventOfCode/Day1"
	"AdventOfCode/Day2"
	"AdventOfCode/Day3"
	"AdventOfCode/Day4"

	//"AdventOfCode/Day5"
	"AdventOfCode/Day10"
	"AdventOfCode/Day11"
	"AdventOfCode/Day6"
	"AdventOfCode/Day7"
	"AdventOfCode/Day8"
	"AdventOfCode/Day9"

	//"AdventOfCode/Day12"
	"AdventOfCode/Day13"
	"AdventOfCode/Day14"
	"AdventOfCode/Day15"

	"fmt"
)

func main() {

	results := [][2]int{
		Day1.Day1(),
		Day2.Day2(),
		Day3.Day3(),
		Day4.Day4(),
		{322500873, 108956227}, //Day5.Day5(),
		Day6.Day6(),
		Day7.Day7(),
		Day8.Day8(),
		Day9.Day9(),
		Day10.Day10(),
		Day11.Day11(),
		{7047, 0}, //Day12.Day12(),
		Day13.Day13(),
		Day14.Day14(),
		Day15.Day15(),
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
