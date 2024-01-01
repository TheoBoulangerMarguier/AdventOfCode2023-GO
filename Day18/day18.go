/*
--- Day 18: Lavaduct Lagoon ---
Thanks to your efforts, the machine parts factory is one of the first factories up and running since the lavafall came back.
However, to catch up with the large backlog of parts requests, the factory will also need a large supply of lava for a while;
the Elves have already started creating a large lagoon nearby for this purpose.

However, they aren't sure the lagoon will be big enough; they've asked you to take a look at the dig plan (your puzzle input). For example:

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
clearing full 1 meter cubes as they go. The directions are given as seen from above, so if "up" were north, then "right"
would be east, and so on. Each trench is also listed with the color that the
edge of the trench should be painted as an RGB hexadecimal color code.

When viewed from above, the above example dig plan would result in the following loop of trench (#)
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
However, this is just the edge of the lagoon; the next step is to dig out the interior so that it is one meter deep as well:

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
While the interior is dug out, the edges are also painted according to the color codes in the dig plan.

The Elves are concerned the lagoon won't be large enough; if they follow their dig plan, how many cubic meters of lava could it hold?
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
		d18p2(),
	}
}

type Instruction struct {
	dir   string
	len   int
	color string
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
		l, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		i := Instruction{
			dir:   split[0],
			len:   l,
			color: strings.Split(strings.Split(split[2], "(")[1], ")")[0],
		}
		instructions = append(instructions, i)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return instructions
}

func d18p1(instructions []Instruction) int {
	coordinates := make([]Point, len(instructions)+1)
	x, y := 0, 0
	sum := 0
	for i, instruction := range instructions {
		coordinates[i] = Point{x, y}
		if instruction.dir == "R" {
			x += instruction.len
		} else if instruction.dir == "L" {
			x -= instruction.len
		} else if instruction.dir == "U" {
			y += instruction.len
		} else if instruction.dir == "D" {
			y -= instruction.len
		}
		sum += instruction.len
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

func d18p2() int {
	return 0
}
