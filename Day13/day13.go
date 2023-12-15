/*
--- Day 13: Point of Incidence ---
With your help, the hot springs team locates an appropriate spring which launches
you neatly and precisely up to the edge of Lava Island.

There's just one problem: you don't see any lava.

You do see a lot of ash and igneous rock; there are even what look like gray
mountains scattered around. After a while, you make your way to a nearby cluster
of mountains only to discover that the valley between them is completely full
of large mirrors. Most of the mirrors seem to be aligned in a consistent way;
perhaps you should head in that direction?

As you move through the valley of mirrors, you find that several of them have
fallen from the large metal frames keeping them in place. The mirrors are extremely
flat and shiny, and many of the fallen mirrors have lodged into the ash at
strange angles. Because the terrain is all one color, it's hard to tell where
it's safe to walk or where you're about to run into a mirror.

You note down the patterns of ash (.) and rocks (#) that you see as you walk
(your puzzle input); perhaps by carefully analyzing these patterns,
you can figure out where the mirrors are!

For example:

#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
To find the reflection in each pattern, you need to find a perfect
reflection across either a horizontal line between two rows or across a vertical
line between two columns.

In the first pattern, the reflection is across a vertical line between two columns;
arrows on each of the two columns point at the line between the columns:

123456789
    ><
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
    ><
123456789
In this pattern, the line of reflection is the vertical line between columns
5 and 6. Because the vertical line is not perfectly in the middle of the pattern,
part of the pattern (column 1) has nowhere to reflect onto and can be ignored;
every other column has a reflected column within the pattern and must match exactly:
column 2 matches column 9, column 3 matches 8, 4 matches 7, and 5 matches 6.

The second pattern reflects across a horizontal line instead:

1 #...##..# 1
2 #....#..# 2
3 ..##..### 3
4v#####.##.v4
5^#####.##.^5
6 ..##..### 6
7 #....#..# 7

This pattern reflects across the horizontal line between rows 4 and 5. Row 1 would
reflect with a hypothetical row 8, but since that's not in the pattern, row 1
doesn't need to match anything. The remaining rows match: row 2 matches row 7, row 3
matches row 6, and row 4 matches row 5.

To summarize your pattern notes, add up the number of columns to the left of
each vertical line of reflection; to that, also add 100 multiplied by the
number of rows above each horizontal line of reflection. In the above example,
the first pattern's vertical line has 5 columns to its left and the second
pattern's horizontal line has 4 rows above it, a total of 405.

Find the line of reflection in each of the patterns in your notes. What number
do you get after summarizing all of your notes?

--- Part Two ---
You resume walking through the valley of mirrors and - SMACK! - run directly into one.
Hopefully nobody was watching, because that must have been pretty embarrassing.

Upon closer inspection, you discover that every mirror has exactly one smudge:
exactly one . or # should be the opposite type.

In each pattern, you'll need to locate and fix the smudge that causes a different
reflection line to be valid. (The old reflection line won't necessarily continue
being valid after the smudge is fixed.)

Here's the above example again:

#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

The first pattern's smudge is in the top-left corner.
If the top-left # were instead ., it would have a different,
horizontal line of reflection:

1 ..##..##. 1
2 ..#.##.#. 2
3v##......#v3
4^##......#^4
5 ..#.##.#. 5
6 ..##..##. 6
7 #.#.##.#. 7

With the smudge in the top-left corner repaired,
a new horizontal line of reflection between rows 3 and 4 now exists.
Row 7 has no corresponding reflected row and can be ignored,
but every other row matches exactly: row 1 matches row 6, row 2 matches row 5,
and row 3 matches row 4.

In the second pattern,
the smudge can be fixed by changing the fifth symbol on row 2 from . to #:

1v#...##..#v1
2^#...##..#^2
3 ..##..### 3
4 #####.##. 4
5 #####.##. 5
6 ..##..### 6
7 #....#..# 7

Now, the pattern has a different horizontal line of reflection between rows 1 and 2.

Summarize your notes as before, but instead use the new different reflection lines.
In this example, the first pattern's new horizontal line has 3 rows above it and
the second pattern's new horizontal line has 1 row above it,
summarizing to the value 400.

In each pattern, fix the smudge and find the different line of reflection.
What number do you get after summarizing the new reflection line in each pattern in your notes?

*/

package Day13

import (
	"bufio"
	"log"
	"math"
	"os"
)

func Day13() [2]int {
	return [2]int{
		d13p1(),
		d13p2(),
	}
}

type Puzzle struct {
	cells     [][]rune
	horitzSym []int
	vertSym   []int
}

func loadDataFromInput(path string) []Puzzle {
	file, err := os.Open(path)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	puzzles := []Puzzle{}
	puzzle := Puzzle{}
	inPuzzle := false

	for scanner.Scan() {
		if scanner.Text() == "" {
			if inPuzzle {
				puzzles = append(puzzles, puzzle)
				inPuzzle = false
				puzzle = Puzzle{}
			}
		} else {
			line := []rune(scanner.Text())
			puzzle.cells = append(puzzle.cells, line)

			if !inPuzzle {
				inPuzzle = true
			}
		}
	}

	if inPuzzle {
		puzzles = append(puzzles, puzzle)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return puzzles
}

func d13p1() int {
	puzzles := loadDataFromInput("./Day13/Ressources/day13_input.txt")
	total := 0
	for i := 0; i < len(puzzles); i++ {
		_, r := calculateFirstReflection(puzzles[i])
		total += r
	}
	return total
}

func d13p2() int {
	puzzles := loadDataFromInput("./Day13/Ressources/day13_input.txt")
	total := 0
	for i := 0; i < len(puzzles); i++ {
		//get V1 of puzzle
		p, _ := calculateFirstReflection(puzzles[i])
		foundSmudge := false
		for y := 0; y < len(p.cells); y++ {
			for x := 0; x < len(p.cells[y]); x++ {

				//make a copy of the puzzle
				cellsUnSmudged := make([][]rune, len(p.cells))
				for i := range p.cells {
					cellsUnSmudged[i] = make([]rune, len(p.cells[i]))
					copy(cellsUnSmudged[i], p.cells[i])
				}

				//swap char after char
				if cellsUnSmudged[y][x] == '.' {
					cellsUnSmudged[y][x] = '#'
				} else {
					cellsUnSmudged[y][x] = '.'
				}

				//get new sim lines
				np, _ := calculateFirstReflection(Puzzle{cells: cellsUnSmudged})

				//remove the old sym line
				dH := DiffSlice(np.horitzSym, p.horitzSym)
				dV := DiffSlice(np.vertSym, p.vertSym)

				if len(dH) != 0 || len(dV) != 0 {
					p.horitzSym = dH
					p.vertSym = dV
					foundSmudge = true
					break
				}
			}
			if foundSmudge {
				if len(p.horitzSym) != 0 {
					total += (100 * p.horitzSym[0])
				}

				if len(p.vertSym) != 0 {
					total += p.vertSym[0]
				}
				break
			}
		}
	}

	return total
}

func calculateFirstReflection(puzzle Puzzle) (Puzzle, int) {
	//horiz sim
	for y := 0; y < len(puzzle.cells)-1; y++ {
		symtest := true
		for x := 0; x < len(puzzle.cells[y]); x++ {
			min := math.Min(float64(y), float64(len(puzzle.cells)-2-y))
			for lookup := 0; lookup <= int(min); lookup++ {
				if puzzle.cells[y-lookup][x] != puzzle.cells[y+lookup+1][x] {
					symtest = false
					break
				}
			}
			if !symtest {
				break
			}
		}
		if symtest {
			puzzle.horitzSym = append(puzzle.horitzSym, y+1)
		}
	}

	//didn't find a horitz sym, need to check for vert sym now
	for x := 0; x < len(puzzle.cells[0])-1; x++ {
		symtest := true
		for y := 0; y < len(puzzle.cells); y++ {
			min := math.Min(float64(x), float64(len(puzzle.cells[0])-2-x))
			for lookup := 0; lookup <= int(min); lookup++ {
				if puzzle.cells[y][x-lookup] != puzzle.cells[y][x+lookup+1] {
					symtest = false
					break
				}
			}
			if !symtest {
				break
			}
		}
		if symtest {
			puzzle.vertSym = append(puzzle.vertSym, x+1)
		}
	}

	total := 0

	if len(puzzle.horitzSym) == 1 {
		total += (100 * puzzle.horitzSym[0])
	}

	if len(puzzle.vertSym) == 1 {
		total += puzzle.vertSym[0]
	}

	return puzzle, total
}

func DiffSlice(A, B []int) []int {
	exists := make(map[int]bool)

	for _, num := range B {
		exists[num] = true
	}

	var diff []int

	for _, num := range A {
		if !exists[num] {
			diff = append(diff, num)
		}
	}
	return diff
}
