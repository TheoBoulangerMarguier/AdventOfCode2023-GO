/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola
lift will take you up to the water source, but this is as far as he can bring
you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem:
they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of
surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working
right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from
the engine, but nobody can figure out which one. If you can add up all
the part numbers in the engine schematic, it should be easy to work out
which part is missing.

The engine schematic (your puzzle input) consists of a visual representation
of the engine. There are lots of numbers and symbols you don't really
understand, but apparently any number adjacent to a symbol, even diagonally,
is a "part number" and should be included in your sum. (Periods (.)
do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not
adjacent to a symbol: 114 (top right) and 58 (middle right).
Every other number is adjacent to a symbol and so is a part number;
their sum is 4361.

Of course, the actual engine schematic is much larger.
What is the sum of all of the part numbers in the engine schematic?


--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine
 springs to life, you jump in the closest gondola, finally ready to ascend to
 the water source.

You don't seem to be going very fast, though. Maybe something is still wrong?
Fortunately, the gondola has a phone labeled "help", so you pick it up and
the engineer answers.

Before you can explain the situation, she suggests that you look out the window.
There stands the engineer, holding a phone in one hand and waving with
the other. You're going so slowly that you haven't even left the station.
You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine
is wrong. A gear is any * symbol that is adjacent to exactly two part numbers.
Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so
that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

package Day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Day3() [2]int {
	return [2]int{
		d3p1(),
		d3p2(),
	}
}

// check if a specific position contains a character different from '.' and not a digit
func evaluatePos(slice [][]rune, predicate bool, x int, y int) bool {
	if predicate {
		r := slice[x][y]
		if r != '.' && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// search if the array all the digit composing a number based on input coordinates
func getFullNumber(slice2D [][]rune, x int, y int) int {
	output := -1

	//early return in case of wrong data input
	if len(slice2D) == 0 {
		return output
	}
	if x < 0 || x > len(slice2D) {
		return output
	}
	if len(slice2D[x]) == 0 {
		return output
	}
	if !unicode.IsDigit(slice2D[x][y]) {
		return output
	}

	sNumber := string(slice2D[x][y])

	//walk left
	ny := y - 1
	for ny >= 0 {
		if unicode.IsDigit(slice2D[x][ny]) {
			sNumber = string(slice2D[x][ny]) + sNumber
		} else {
			break
		}
		ny--
	}

	//walk right
	ny = y + 1
	for ny < len(slice2D) {
		if unicode.IsDigit(slice2D[x][ny]) {
			sNumber += string(slice2D[x][ny])
		} else {
			break
		}
		ny++
	}

	if n, err := strconv.Atoi(sNumber); err == nil {
		return n
	} else {
		log.Fatal(err)
	}
	return output
}

func d3p1() int {
	//step 0 open the file
	file, err := os.Open("./Day3/Ressources/day3_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	//step 1 get document size in order to initialize the 2D array
	var slice2D [][]rune

	for scanner.Scan() {
		slice2D = append(slice2D, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//step 2 walk over and find numbers and their surroundings
	sum := 0
	currentNumber := ""
	valid := false

	for x := 0; x < len(slice2D); x++ {
		for y := 0; y < len(slice2D[x]); y++ {
			if unicode.IsDigit(slice2D[x][y]) {
				currentNumber += string(slice2D[x][y])

				//check for non '.' and non digit character in the 8 cells around the current digit
				//check left
				hasLeft := x > 0
				hasRight := x < len(slice2D)-1
				hasUp := y < len(slice2D[x])-1
				hasDown := y > 0

				valid = valid || evaluatePos(slice2D, hasLeft, x-1, y)
				valid = valid || evaluatePos(slice2D, hasRight, x+1, y)
				valid = valid || evaluatePos(slice2D, hasDown, x, y-1)
				valid = valid || evaluatePos(slice2D, hasUp, x, y+1)
				valid = valid || evaluatePos(slice2D, hasLeft && hasUp, x-1, y+1)
				valid = valid || evaluatePos(slice2D, hasRight && hasUp, x+1, y+1)
				valid = valid || evaluatePos(slice2D, hasLeft && hasDown, x-1, y-1)
				valid = valid || evaluatePos(slice2D, hasRight && hasDown, x+1, y-1)

			} else {
				//close the current number and evaluate validity to add to sum
				if currentNumber != "" && valid {
					if i, err := strconv.Atoi(currentNumber); err == nil {
						sum += i
					}
				}
				currentNumber = ""
				valid = false
			}

		}
		//close the line if any numbers is still pending
		if currentNumber != "" && valid {
			if i, err := strconv.Atoi(currentNumber); err == nil {
				sum += i
			}
		}
		currentNumber = ""
		valid = false
	}

	return sum
}

func d3p2() int {
	//step 0 open the file
	file, err := os.Open("./Day3/Ressources/day3_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	//step 1 get document size in order to initialize the 2D array
	var slice2D [][]rune

	for scanner.Scan() {
		slice2D = append(slice2D, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//will be the output of the function, sum of all gear ratio found
	sum := 0

	//step 2 walk over and find numbers and their surroundings
	for x := 0; x < len(slice2D); x++ {
		for y := 0; y < len(slice2D[x]); y++ {

			//find a gear
			if slice2D[x][y] == '*' {
				//get the 8 cell around it, for simplicity we will consider out
				//of bounds and non-digit as '.' and ignore them
				neighbour := [3][3]rune{
					{'.', '.', '.'},
					{'.', '*', '.'},
					{'.', '.', '.'},
				}

				for offsetX := -1; offsetX <= 1; offsetX++ {
					for offsetY := -1; offsetY <= 1; offsetY++ {
						if offsetX == 0 && offsetY == 0 {
							continue
						}
						//checking if the offset is in bound
						inBounds := x+offsetX < len(slice2D[x])
						inBounds = inBounds && x+offsetX >= 0
						inBounds = inBounds && y+offsetY < len(slice2D)
						inBounds = inBounds && y+offsetY >= 0

						//remap neighbour coordinates from -1->1 to 0->2 by adding 1
						if inBounds {
							neighbour[offsetX+1][offsetY+1] = slice2D[x+offsetX][y+offsetY]
						}
					}
				}

				//count how many possible number we have
				countNumbers := 0
				isOverNumber := false
				gearRatio := 1

				//walk through neighbours to find numbers and number count
				for nx := 0; nx < len(neighbour); nx++ {
					for ny := 0; ny < len(neighbour[nx]); ny++ {
						if !isOverNumber && unicode.IsDigit(neighbour[nx][ny]) {
							isOverNumber = true
							//get the full number from the current digit
							//(-1 on the neighbourt pos to convert to global pos)
							gearRatio *= getFullNumber(slice2D, x+nx-1, y+ny-1)
						} else if isOverNumber && !unicode.IsDigit(neighbour[nx][ny]) {
							isOverNumber = false
							countNumbers++
						}
					}
					if isOverNumber {
						isOverNumber = false
						countNumbers++
					}
				}

				//cumul the gear ration for gear that have exactly 2 surroundign numbers
				if countNumbers == 2 {
					sum += gearRatio
				}
			}
		}
	}
	return sum
}
