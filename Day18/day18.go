/*
--- Day 18: Lavaduct Lagoon ---
Thanks to your efforts, the machine parts factory is one of the first factories
up and running since the lavafall came back.
However, to catch up with the large backlog of parts requests, the factory will
also need a large supply of lava for a while;
the Elves have already started creating a large lagoon nearby for this purpose.

However, they aren't sure the lagoon will be big enough; they've asked you to
take a look at the dig plan (your puzzle input). For example:

R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)

The digger starts in a 1 meter cube hole in the ground.
They then dig the specified number of meters up (U), down (D), left (L), or right (R),
clearing full 1 meter cubes as they go. The directions are given as seen from
above, so if "up" were north, then "right"
would be east, and so on. Each trench is also listed with the color that the
edge of the trench should be painted as an RGB hexadecimal color code.

When viewed from above, the above example dig plan would result in the following
loop of trench (#)
having been dug out from otherwise ground-level terrain (.):

#######
#.....#
###...#
..#...#
..#...#
###.###
#...#..
##..###
.#....#
.######

At this point, the trench could contain 38 cubic meters of lava.
However, this is just the edge of the lagoon; the next step is to dig out the
interior so that it is one meter deep as well:

#######
#######
#######
..#####
..#####
#######
#####..
#######
.######
.######

Now, the lagoon can contain a much more respectable 62 cubic meters of lava.
While the interior is dug out, the edges are also painted according to the color
codes in the dig plan.

The Elves are concerned the lagoon won't be large enough; if they follow their
dig plan, how many cubic meters of lava could it hold?

--- Part Two ---
The Elves were right to be concerned; the planned lagoon would be much too small.

After a few minutes, someone realizes what happened; someone swapped the color
and instruction parameters when producing the dig plan. They don't have time to
fix the bug; one of them asks if you can extract the correct instructions from
the hexadecimal codes.

Each hexadecimal code is six hexadecimal digits long. The first five hexadecimal
digits encode the distance in meters as a five-digit hexadecimal number.
The last hexadecimal digit encodes the direction to dig: 0 means R, 1
means D, 2 means L, and 3 means U.

So, in the above example, the hexadecimal codes can be converted into the true
instructions:

#70c710 = R 461937
#0dc571 = D 56407
#5713f0 = R 356671
#d2c081 = D 863240
#59c680 = R 367720
#411b91 = D 266681
#8ceee2 = L 577262
#caa173 = U 829975
#1b58a2 = L 112010
#caa171 = D 829975
#7807d2 = L 491645
#a77fa3 = U 686074
#015232 = L 5411
#7a21e3 = U 500254

Digging out this loop and its interior produces a lagoon that can hold an impressive
952408144115 cubic meters of lava.

Convert the hexadecimal color codes into the correct instructions;
if the Elves follow this new dig plan, how many cubic meters of lava
could the lagoon hold?
*/

package Day18

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day18() [2]int {
	i := loadData()
	return [2]int{
		d18p1(i),
		d18p2(i),
	}
}

type Instruction struct {
	dir    string
	len    int
	hexDir string
	hexLen int
}

type Point struct {
	x, y int
}

func loadData() []Instruction {
	file, err := os.Open("./Day18/Ressources/day18_input.txt")
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)
	instructions := []Instruction{}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		//get strandard len (p1)
		l, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		//convert hexColor hex to dir and len (p2)
		hexColor := strings.Split(strings.Split(split[2], "(")[1], ")")[0]
		hexDirRaw := string([]rune(hexColor)[len(hexColor)-1])
		hexDirInt, err := strconv.Atoi(hexDirRaw)
		if err != nil {
			log.Fatal(err)
		}
		hexDir := intToStringDir(hexDirInt)
		hexLenRaw := string([]rune(hexColor)[1 : len(hexColor)-1])
		hexLenInt, err := strconv.ParseInt(hexLenRaw, 16, 64)
		if err != nil {
			log.Fatal(err)
		}

		//create the new instruction
		newInstruction := Instruction{
			dir:    split[0],
			len:    l,
			hexDir: hexDir,
			hexLen: int(hexLenInt),
		}

		instructions = append(instructions, newInstruction)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return instructions
}

func d18p1(instructions []Instruction) int {
	return getAreaPickShoelace(instructions, false)
}

func d18p2(instructions []Instruction) int {
	return getAreaPickShoelace(instructions, true)
}

func intToStringDir(i int) string {
	if i == 0 {
		return "R"
	} else if i == 1 {
		return "D"
	} else if i == 2 {
		return "L"
	} else {
		return "U"
	}
}

func getAreaPickShoelace(input []Instruction, isHex bool) int {
	coordinates := make([]Point, len(input)+1)
	x, y := 0, 0
	sum := 0
	for i, instruction := range input {
		coordinates[i] = Point{x, y}
		var dir string
		var len int

		if isHex {
			dir = instruction.hexDir
			len = instruction.hexLen
		} else {
			dir = instruction.dir
			len = instruction.len
		}

		if dir == "R" {
			x += len
		} else if dir == "L" {
			x -= len
		} else if dir == "U" {
			y += len
		} else if dir == "D" {
			y -= len
		}
		sum += len
	}

	//shoelace theorem
	// A = 0.5*Abs((x1y2 + x2y3 + ...)-(y1x2 + y2x3 + ...))
	sumA := 0
	sumB := 0
	for i := 0; i < len(coordinates)-1; i++ {
		sumA += (coordinates[i].x * coordinates[i+1].y)
		sumB += (coordinates[i].y * coordinates[i+1].x)
	}

	area := math.Abs(float64(sumA)-float64(sumB)) * 0.5

	//Pick's theorem
	// A = i + B*0.5-1
	//where A = area (shoelace result), i = inside amount of points, B = border amount of point
	// we need i = ?
	// we have b = sum, A = area
	// we solve with => i = area - b/2 +1

	total := area - float64(sum)/2 + 1 + float64(sum)
	return int(total)
}
