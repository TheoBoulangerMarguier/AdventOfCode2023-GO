/*
--- Day 22: Sand Slabs ---
Enough sand has fallen; it can finally filter water for Snow Island.

Well, almost.

The sand has been falling as large compacted bricks of sand, piling up to form
an impressive stack here near the edge of Island Island.
In order to make use of the sand to filter water, some of the bricks will need
to be broken apart - nay, disintegrated - back into freely flowing sand.

The stack is tall enough that you'll have to be careful about choosing which
bricks to disintegrate; if you disintegrate the wrong brick, large portions of
the stack could topple, which sounds pretty dangerous.

The Elves responsible for water filtering operations took a snapshot of the
bricks while they were still falling (your puzzle input) which should
let you work out which bricks are safe to disintegrate. For example:

1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
Each line of text in the snapshot represents the position of a single brick at
the time the snapshot was taken. The position is given as
two x,y,z coordinates - one for each end of the brick - separated by a tilde (~).
Each brick is made up of a single straight line of cubes, and the Elves were even
careful to choose a time for the snapshot that had all of the free-falling
bricks at integer positions above the ground, so the whole snapshot is aligned
to a three-dimensional cube grid.

A line like 2,2,2~2,2,2 means that both ends of the brick are at the
same coordinate - in other words, that the brick is a single cube.

Lines like 0,0,10~1,0,10 or 0,0,10~0,1,10 both represent bricks that are
two cubes in volume, both oriented horizontally.
The first brick extends in the x direction, while the second brick extends
in the y direction.

A line like 0,0,1~0,0,10 represents a ten-cube brick which is oriented vertically.
One end of the brick is the cube located at 0,0,1,
while the other end of the brick is located directly above it at 0,0,10.

The ground is at z=0 and is perfectly flat; the lowest z value a brick can have
is therefore 1. So, 5,5,1~5,6,1 and 0,2,1~0,2,5 are both resting on the ground,
but 3,3,2~3,3,3 was above the ground at the time of the snapshot.

Because the snapshot was taken while the bricks were still falling,
some bricks will still be in the air; you'll need to start by figuring out where
they will end up. Bricks are magically stabilized, so they never rotate,
even in weird situations like where a long horizontal brick is only
supported on one end. Two bricks cannot occupy the same position,
so a falling brick will come to rest upon the first other brick it encounters.

Here is the same example again, this time with each brick given
a letter so it can be marked in diagrams:

1,0,1~1,2,1   <- A
0,0,2~2,0,2   <- B
0,2,3~2,2,3   <- C
0,0,4~0,2,4   <- D
2,0,5~2,2,5   <- E
0,1,6~2,1,6   <- F
1,1,8~1,1,9   <- G

At the time of the snapshot, from the side so the x axis goes left to right,
these bricks are arranged like this:

 x
012
.G. 9
.G. 8
... 7
FFF 6
..E 5 z
D.. 4
CCC 3
BBB 2
.A. 1
--- 0

Rotating the perspective 90 degrees so the y axis now goes left to right,
the same bricks are arranged like this:

 y
012
.G. 9
.G. 8
... 7
.F. 6
EEE 5 z
DDD 4
..C 3
B.. 2
AAA 1
--- 0

Once all of the bricks fall downward as far as they can go, the stack looks
like this, where ? means bricks are hidden behind other bricks at that location:

 x
012
.G. 6
.G. 5
FFF 4
D.E 3 z
??? 2
.A. 1
--- 0

Again from the side:

 y
012
.G. 6
.G. 5
.F. 4
??? 3 z
B.C 2
AAA 1
--- 0

Now that all of the bricks have settled,
it becomes easier to tell which bricks are supporting which other bricks:

Brick A is the only brick supporting bricks B and C.
Brick B is one of two bricks supporting brick D and brick E.
Brick C is the other brick supporting brick D and brick E.
Brick D supports brick F.
Brick E also supports brick F.
Brick F supports brick G.
Brick G isn't supporting any bricks.

Your first task is to figure out which bricks are safe to disintegrate.
A brick can be safely disintegrated if, after removing it, no other bricks would
fall further directly downward.
Don't actually disintegrate any bricks - just determine what would happen if,
for each brick, only that brick were disintegrated. Bricks can be disintegrated
even if they're completely surrounded by other bricks; you can squeeze
between bricks if you need to.

In this example, the bricks can be disintegrated as follows:

Brick A cannot be disintegrated safely; if it were disintegrated, bricks B and C would both fall.
Brick B can be disintegrated; the bricks above it (D and E) would still be supported by brick C.
Brick C can be disintegrated; the bricks above it (D and E) would still be supported by brick B.
Brick D can be disintegrated; the brick above it (F) would still be supported by brick E.
Brick E can be disintegrated; the brick above it (F) would still be supported by brick D.
Brick F cannot be disintegrated; the brick above it (G) would fall.
Brick G can be disintegrated; it does not support any other bricks.
So, in this example, 5 bricks can be safely disintegrated.

Figure how the blocks will settle based on the snapshot.
Once they've settled, consider disintegrating a single brick;
how many bricks could be safely chosen as the one to get disintegrated?

--- Part Two ---
Disintegrating bricks one at a time isn't going to be fast enough.
While it might sound dangerous, what you really need is a chain reaction.

You'll need to figure out the best brick to disintegrate.
For each brick, determine how many other bricks would fall if that brick were disintegrated.

Using the same example as above:

Disintegrating brick A would cause all 6 other bricks to fall.
Disintegrating brick F would cause only 1 other brick, G, to fall.
Disintegrating any other brick would cause no other bricks to fall.
So, in this example, the sum of the number of other bricks
that would fall as a result of disintegrating each brick is 7.

For each brick, determine how many other bricks would fall
if that brick were disintegrated.
What is the sum of the number of other bricks that would fall?

*/

package Day22

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// x is left/right, y is forward/backward, z is up/down
type Point3 struct {
	x, y, z int
}

type Brick struct {
	start, end   Point3
	allPos, base []Point3
	supportedBy  []int
	isSupporting []int
}

func Day22() [2]int {
	return [2]int{
		d22p1(),
		d22p2(),
	}
}

func loadData() (map[int]Brick, map[Point3]int, Point3) {
	file, err := os.Open("./Day22/Ressources/day22_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bricks := map[int]Brick{}
	brickID := 1
	bounds := Point3{-1, -1, -1}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newBrick := Brick{}

		//Parse start and end positions
		pos2 := strings.Split(scanner.Text(), "~")
		startEnd := [][]string{
			strings.Split(pos2[0], ","),
			strings.Split(pos2[1], ","),
		}
		for i, pos := range startEnd {
			for j, coord := range pos {
				n, err := strconv.Atoi(coord)
				if err != nil {
					log.Fatal(err)
				}
				if i == 0 {
					if j == 0 {
						newBrick.start.x = n
					} else if j == 1 {
						newBrick.start.y = n
					} else {
						newBrick.start.z = n
					}
				} else {
					if j == 0 {
						newBrick.end.x = n
					} else if j == 1 {
						newBrick.end.y = n
					} else {
						newBrick.end.z = n
					}
				}
			}
		}

		//update bounds
		if newBrick.start.x > bounds.x {
			bounds.x = newBrick.start.x
		}
		if newBrick.end.x > bounds.x {
			bounds.x = newBrick.end.x
		}
		if newBrick.start.y > bounds.y {
			bounds.y = newBrick.start.y
		}
		if newBrick.end.y > bounds.y {
			bounds.y = newBrick.end.y
		}
		if newBrick.start.z > bounds.z {
			bounds.z = newBrick.start.z
		}
		if newBrick.end.z > bounds.z {
			bounds.z = newBrick.end.z
		}

		bricks[brickID] = newBrick
		brickID++
	}

	//create 3D
	grid := map[Point3]int{}
	for x := 0; x < bounds.x; x++ {
		for y := 0; y < bounds.y; y++ {
			for z := 0; z < bounds.z; z++ {
				grid[Point3{x, y, z}] = 0
			}
		}
	}

	//places blocks ref in grid, start pos is allways smaller or equal to end pos
	for key, brick := range bricks {
		counts := Point3{
			brick.end.x - brick.start.x,
			brick.end.y - brick.start.y,
			brick.end.z - brick.start.z,
		}

		//generate positions from start to end and store them in the allPos of the brick
		//store the brick ID in the grid for all occupied positions
		//start is always base
		brick.allPos = append(brick.allPos, brick.start)
		brick.base = append(brick.base, brick.start)
		grid[brick.start] = key

		for x := 1; x <= counts.x; x++ {
			pos := Point3{
				brick.start.x + x,
				brick.start.y,
				brick.start.z,
			}
			brick.allPos = append(brick.allPos, pos)
			brick.base = append(brick.base, pos)
			grid[pos] = key
		}
		for y := 1; y <= counts.y; y++ {
			pos := Point3{
				brick.start.x,
				brick.start.y + y,
				brick.start.z,
			}
			brick.allPos = append(brick.allPos, pos)
			brick.base = append(brick.base, pos)
			grid[pos] = key
		}
		for z := 1; z <= counts.z; z++ {
			pos := Point3{
				brick.start.x,
				brick.start.y,
				brick.start.z + z,
			}
			brick.allPos = append(brick.allPos, pos)
			grid[pos] = key
		}
		grid[brick.end] = key
		bricks[key] = brick
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return bricks, grid, bounds
}

func applyGravity(bricks map[int]Brick, grid map[Point3]int) {
	//apply gravity to the bricks
	for {
		hasMovedThisTurn := false
		for key, brick := range bricks {
			//check if the brick can drop by looking down from all "base" pos
			canDrop := true
			for _, pos := range brick.base {
				if pos.z == 1 {
					canDrop = false
					break //already at ground level
				}
				nextPos := Point3{
					pos.x,
					pos.y,
					pos.z - 1,
				}
				if grid[nextPos] != 0 {
					canDrop = false
					break //level below is already occupied
				}
			}
			//drop by one level and update the grid and position info
			if canDrop {
				hasMovedThisTurn = true
				newAllPos := []Point3{}
				newBases := []Point3{}

				for _, pos := range brick.allPos {
					grid[pos] = 0 //clear grid from old info
					nextPos := Point3{
						pos.x,
						pos.y,
						pos.z - 1,
					}

					newAllPos = append(newAllPos, nextPos)
					isBase, err := utils.SliceContains(brick.base, pos)
					if err != nil {
						log.Fatal(err)
					}
					if isBase {
						newBases = append(newBases, nextPos)
					}
				}

				//replace old pos by new pos
				brick.allPos = newAllPos
				brick.base = newBases
				bricks[key] = brick

				for _, pos := range brick.allPos {
					grid[pos] = key // apply new info to the grid
				}
			}
		}

		if !hasMovedThisTurn {
			break
		}
	}

	//get all the "supporting bricks"
	for key, brick := range bricks {
		for _, pos := range brick.base {
			nextPos := Point3{
				pos.x,
				pos.y,
				pos.z - 1,
			}
			if grid[nextPos] != 0 {
				if ok, err := utils.SliceContains(brick.supportedBy, grid[nextPos]); !ok {
					brick.supportedBy = append(brick.supportedBy, grid[nextPos])
					supportBrick := bricks[grid[nextPos]]
					supportBrick.isSupporting = append(supportBrick.isSupporting, key)
					bricks[grid[nextPos]] = supportBrick
				} else if err != nil {
					log.Fatal(err)
				}
			}
			bricks[key] = brick

		}
	}
}

func d22p1() int {
	bricks, grid, _ := loadData()
	applyGravity(bricks, grid)

	//find how many can be safely disintegrated
	//a brick can be disintegrated if all the suported brick have at least another supporter
	toDestroy := []int{}
	for key, brick1 := range bricks {
		canDestroy := true
		for _, brick2ID := range brick1.isSupporting {
			if len(bricks[brick2ID].supportedBy) < 2 {
				canDestroy = false
			}
		}

		if canDestroy {
			toDestroy = append(toDestroy, key)
		}
	}
	return len(toDestroy)
}

func d22p2() int {
	bricks, grid, _ := loadData()
	applyGravity(bricks, grid)

	sum := 0

	//we will look at each bricks to see what hapens if we destroy them
	for id0, _ := range bricks {
		areGoingTofall := []int{id0}
		//we will repeat the process below for each brick known as going to fall
		for i := 0; i < len(areGoingTofall); i++ {
			//we want to confirm if all the bricks suported by the curretn ID will fall
			for _, id1 := range bricks[areGoingTofall[i]].isSupporting {
				willFall := true
				//for that we need to know if all of the supporting bricks(id2)
				//of the suported id1 brick are part of the "going to fall" list
				//otherwise it mean they are supported by a stable brick
				for _, id2 := range bricks[id1].supportedBy {
					if ok, err := utils.SliceContains(areGoingTofall, id2); err != nil {
						log.Fatal(err)

					} else if !ok {
						willFall = false
						break
					}
				}
				//once confirmed we can add this brick to the list so that we continue to check going up
				if willFall {
					if ok, err := utils.SliceContains(areGoingTofall, id1); err != nil {
						log.Fatal(err)
					} else if !ok {
						areGoingTofall = append(areGoingTofall, id1)
					}

				}
			}
		}

		if len(areGoingTofall) > 1 {
			sum += len(areGoingTofall)
			sum-- //removing the extra count of the first brick being destroyed and not falling
		}
	}

	return sum
}

//BELOW FUNCTION CAN BE USED TO VISUALIZE THE GRID IN A SIMILAR WAY AS THE EXAMPLES
/*
	func printGridViewX(grid map[Point3]int, bounds Point3) {
		fmt.Println("View X:")
		for z := bounds.z; z >= 0; z-- {
			fmt.Print(z, "  ")
			for x := 0; x <= bounds.x; x++ {
				inRow := []int{}
				for y := 0; y <= bounds.y; y++ {
					id := grid[Point3{x, y, z}]
					ok, err := utils.SliceContains(inRow, id)
					if err != nil {
						log.Fatal(err)
					}

					if id != 0 && !ok {
						inRow = append(inRow, id)
					}
				}
				if z == 0 {
					fmt.Print("-")
				} else if len(inRow) == 1 {
					fmt.Print(inRow[0])
				} else if len(inRow) > 1 {
					fmt.Print("?")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	func printGridViewY(grid map[Point3]int, bounds Point3) {
		fmt.Println("View Y:")
		for z := bounds.z; z >= 0; z-- {
			fmt.Print(z, "  ")
			for y := 0; y <= bounds.y; y++ {
				inRow := []int{}
				for x := 0; x <= bounds.x; x++ {
					id := grid[Point3{x, y, z}]
					ok, err := utils.SliceContains(inRow, id)
					if err != nil {
						log.Fatal(err)
					}

					if id != 0 && !ok {
						inRow = append(inRow, id)
					}
				}
				if z == 0 {
					fmt.Print("-")
				} else if len(inRow) == 1 {
					fmt.Print(inRow[0])
				} else if len(inRow) > 1 {
					fmt.Print("?")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
func printSupport(bricks map[int]Brick) {
	for i := 1; i <= len(bricks); i++ {
		fmt.Println("brick", i, "is supporting:", bricks[i].isSupporting, "and is supported by:", bricks[i].supportedBy)
	}
}
*/
