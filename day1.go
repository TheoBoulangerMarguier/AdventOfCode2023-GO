/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to
take a look. The Elves have even given you a map; on it, they've used stars
 to mark the top fifty locations that are likely to be having problems.
You've been doing this long enough to know that to restore snow operations,
you need to check all fifty stars by December 25th.
Collect stars by solving puzzles. Two puzzles will be made available
on each day in the Advent calendar; the second puzzle is unlocked when you
complete the first. Each puzzle grants one star. Good luck!
You try to ask why they can't just use a weather machine ("not powerful enough")
and where they're even sending you ("the sky") and why your map looks mostly
blank ("you sure ask a lot of questions") and hang on did you just say the sky
("of course, where do you think snow comes from") when you realize
that the Elves are already loading you into a trebuchet
("please hold still, we need to strap you in").
As they're making the final adjustments, they discover that their calibration
document (your puzzle input) has been amended by a very young Elf who was
apparently just excited to show off her art skills. Consequently, the Elves
are having trouble reading the values on the document.
The newly-improved calibration document consists of lines of text; each line
originally contained a specific calibration value that the Elves
now need to recover. On each line, the calibration value can be found by
combining the first digit and the last digit (in that order)
to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are
12, 38, 15, and 77. Adding these together produces 142.
Consider your entire calibration document. What is the sum of all of the
calibration values?

Your puzzle answer was 54916.

The first half of this puzzle is complete! It provides one gold star: *


--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are
actually spelled out with letters: one, two, three, four, five, six, seven,
eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and
last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76.
Adding these together produces 281.

What is the sum of all of the calibration values?
*/

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Day1() [2]int {
	return [2]int{
		d1p1(),
		d1p2(),
	}
}

func d1p1() int {
	//Step 1 open and scan the file
	file, err := os.Open("./Ressources/day1_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		//find first and last digit in the string
		chars := []rune(scanner.Text())
		firstDigit := rune(0)
		lastDigit := rune(0)

		for i := 0; i < len(chars); i++ {
			if unicode.IsDigit(chars[i]) {
				if firstDigit == rune(0) {
					firstDigit = chars[i]
					lastDigit = chars[i]

				} else {
					lastDigit = chars[i]
				}
			}
		}

		//convert the 2 single digits strings into one number
		doubleDigit, err := strconv.Atoi(string(firstDigit) + string(lastDigit))

		if err != nil {
			log.Fatal(err)
		}

		sum += doubleDigit
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func d1p2() int {
	lookup := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	//step 1 parse the file and scan it
	file, err := os.Open("./Ressources/day1_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sum := 0

	//look for text matching lookup table or digit
	//convert the textNumbers to digit
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := []rune(scanner.Text())
		firstDigit := rune(0)
		lastDigit := rune(0)

		for i := 0; i < len(chars); i++ {
			if unicode.IsDigit(chars[i]) {
				if firstDigit == rune(0) {
					firstDigit = chars[i]
					lastDigit = chars[i]

				} else {
					lastDigit = chars[i]
				}
			} else {
				for key, value := range lookup {
					if i+len(key) <= len(chars) {
						if string(chars[i:i+len(key)]) == key {
							if firstDigit == rune(0) {
								firstDigit = value
								lastDigit = value

							} else {
								lastDigit = value
							}
						}
					}
				}
			}
		}

		doubleDigit, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		if err != nil {
			log.Fatal(err)
		}

		sum += doubleDigit
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
