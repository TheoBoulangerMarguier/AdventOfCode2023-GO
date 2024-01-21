/* --- Day 24: Never Tell Me The Odds ---
It seems like something is going wrong with the snow-making process. Instead of
forming snow, the water that's been absorbed into the air seems to be forming hail!

Maybe there's something you can do to break up the Vector3s?

Due to strong, probably-magical winds, the Vector3s are all flying through the
air in perfectly linear trajectories. You make a note of each Vector3's
position and velocity (your puzzle input). For example:

19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3

Each line of text corresponds to the position and velocity of a single Vector3.
The positions indicate where the Vector3s are right now (at time 0).
The velocities are constant and indicate exactly how far each Vector3 will
move in one nanosecond.

Each line of text uses the format px py pz @ vx vy vz. For instance,
the Vector3 specified by 20, 19, 15 @ 1, -5, -3
has initial X position 20, Y position 19, Z position 15, X velocity 1, Y velocity -5, and Z velocity -3.
After one nanosecond, the Vector3 would be at 21, 14, 12.

Perhaps you won't have to do anything.
How likely are the Vector3s to collide with each other and smash into tiny ice crystals?

To estimate this, consider only the X and Y axes; ignore the Z axis.
Looking forward in time, how many of the Vector3s' paths will intersect within a test area?
(The Vector3s themselves don't have to collide, just test for intersections between the paths they will trace.)

In this example, look for intersections that happen with an X and Y position each at least 7 and at most 27;
in your actual data, you'll need to check a much larger test area.
Comparing all pairs of Vector3s' future paths produces the following results:

Vector3 A: 19, 13, 30 @ -2, 1, -2
Vector3 B: 18, 19, 22 @ -1, -1, -2
Vector3s' paths will cross inside the test area (at x=14.333, y=15.333).

Vector3 A: 19, 13, 30 @ -2, 1, -2
Vector3 B: 20, 25, 34 @ -2, -2, -4
Vector3s' paths will cross inside the test area (at x=11.667, y=16.667).

Vector3 A: 19, 13, 30 @ -2, 1, -2
Vector3 B: 12, 31, 28 @ -1, -2, -1
Vector3s' paths will cross outside the test area (at x=6.2, y=19.4).

Vector3 A: 19, 13, 30 @ -2, 1, -2
Vector3 B: 20, 19, 15 @ 1, -5, -3
Vector3s' paths crossed in the past for Vector3 A.

Vector3 A: 18, 19, 22 @ -1, -1, -2
Vector3 B: 20, 25, 34 @ -2, -2, -4
Vector3s' paths are parallel; they never intersect.

Vector3 A: 18, 19, 22 @ -1, -1, -2
Vector3 B: 12, 31, 28 @ -1, -2, -1
Vector3s' paths will cross outside the test area (at x=-6, y=-5).

Vector3 A: 18, 19, 22 @ -1, -1, -2
Vector3 B: 20, 19, 15 @ 1, -5, -3
Vector3s' paths crossed in the past for both Vector3s.

Vector3 A: 20, 25, 34 @ -2, -2, -4
Vector3 B: 12, 31, 28 @ -1, -2, -1
Vector3s' paths will cross outside the test area (at x=-2, y=3).

Vector3 A: 20, 25, 34 @ -2, -2, -4
Vector3 B: 20, 19, 15 @ 1, -5, -3
Vector3s' paths crossed in the past for Vector3 B.

Vector3 A: 12, 31, 28 @ -1, -2, -1
Vector3 B: 20, 19, 15 @ 1, -5, -3
Vector3s' paths crossed in the past for both Vector3s.

So, in this example, 2 Vector3s' future paths cross inside the boundaries of the test area.

However, you'll need to search a much larger test area if you want to see if any Vector3s might collide.
Look for intersections that happen with an X and Y position each at least 200000000000000 and at most 400000000000000.
Disregard the Z axis entirely.

Considering only the X and Y axes, check all pairs of Vector3s' future paths for intersections.
How many of these intersections occur within the test area?

--- Part Two ---
Upon further analysis, it doesn't seem like any Vector3s will naturally collide.
It's up to you to fix that!

You find a rock on the ground nearby. While it seems extremely unlikely,
if you throw it just right, you should be able to hit every Vector3 in a single throw!

You can use the probably-magical winds to reach any integer position you like and
to propel the rock at any integer velocity. Now including the Z axis in your
calculations, if you throw the rock at time 0, where do you need to be so that
the rock perfectly collides with every Vector3? Due to probably-magical inertia,
the rock won't slow down or change direction when it collides with a Vector3.

In the example above, you can achieve this by moving to position 24, 13, 10 and
throwing the rock at velocity -3, 1, 2. If you do this, you will hit every Vector3 as follows:

Vector3: 19, 13, 30 @ -2, 1, -2
Collision time: 5
Collision position: 9, 18, 20

Vector3: 18, 19, 22 @ -1, -1, -2
Collision time: 3
Collision position: 15, 16, 16

Vector3: 20, 25, 34 @ -2, -2, -4
Collision time: 4
Collision position: 12, 17, 18

Vector3: 12, 31, 28 @ -1, -2, -1
Collision time: 6
Collision position: 6, 19, 22

Vector3: 20, 19, 15 @ 1, -5, -3
Collision time: 1
Collision position: 21, 14, 12

Above, each Vector3 is identified by its initial position and its velocity.
Then, the time and position of that Vector3's collision with your rock are given.

After 1 nanosecond, the rock has exactly the same position as one of the Vector3s,
obliterating it into ice dust! Another Vector3 is smashed to bits two nanoseconds after that.
After a total of 6 nanoseconds, all of the Vector3s have been destroyed.

So, at time 0, the rock needs to be at X position 24, Y position 13, and Z position 10.
Adding these three coordinates together produces 47. (Don't add any coordinates from the rock's velocity.)

Determine the exact position and velocity the rock needs to have at time 0 so
that it perfectly collides with every Vector3.
What do you get if you add up the X, Y, and Z coordinates of that initial position?
*/

package Day24

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day24() [2]int {
	return [2]int{
		d24p1(),
		d24p2(),
	}
}

func d24p1() int {
	datas := loadData("./Day24/Ressources/day24_input.txt")
	sum := 0

	for i := 0; i < len(datas)-1; i++ {
		for j := i + 1; j < len(datas); j++ {

			vA := datas[i]
			vB := datas[j]
			result := checkPositionInTestArea(200000000000000, 400000000000000, vA, vB)
			if result {
				sum++
			}
		}
	}

	return sum
}

func d24p2() int {
	return 0
}

type Point3 struct {
	x, y, z float64
}

type Vector3 struct {
	pos, dir Point3
}

func dot(a, b Point3, dim int) float64 {

	if dim == 2 {
		return a.x*b.x + a.y*b.y
	}

	if dim == 3 {
		return a.x*b.x + a.y*b.y + a.z*b.z
	}

	return math.Inf(1)
}

func cross(a, b Point3) Point3 {
	return Point3{
		a.y*b.z - b.y*a.z,
		a.z*b.x - b.z*a.x,
		a.x*b.y - b.x*a.y,
	}
}

//_____________________________________________________________________________
//______________________________________PART 1_________________________________
//_____________________________________________________________________________

func loadData(path string) []Vector3 {
	//open text file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	//scan through the file with scanner
	scanner := bufio.NewScanner(file)
	datas := []Vector3{}
	for scanner.Scan() {
		//Each line of text uses the format px py pz @ vx vy vz
		twoPart := strings.Split(scanner.Text(), " @ ")
		pos := strings.Split(twoPart[0], ", ")
		dir := strings.Split(twoPart[1], ", ")

		pX, errPx := strconv.Atoi(pos[0])
		pY, errPy := strconv.Atoi(pos[1])
		pZ, errPz := strconv.Atoi(pos[2])
		dX, errDx := strconv.Atoi(dir[0])
		dY, errDy := strconv.Atoi(dir[1])
		dZ, errDz := strconv.Atoi(dir[2])

		if errPx != nil || errPy != nil || errPz != nil {
			log.Fatal("[loadData] error while parsing pos to int (x,y,z):", errPx, errPy, errPz)
		}

		if errDx != nil || errDy != nil || errDz != nil {
			log.Fatal("[loadData] error while parsing dir to int (x,y,z):", errDx, errDy, errDz)
		}

		vector := Vector3{
			pos: Point3{
				x: float64(pX),
				y: float64(pY),
				z: float64(pZ),
			},
			dir: Point3{
				x: float64(dX),
				y: float64(dY),
				z: float64(dZ),
			},
		}

		datas = append(datas, vector)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return datas
}

func findLineEquation(v3 Vector3) (float64, float64) {
	A := v3.pos
	B := Point3{
		x: v3.pos.x + v3.dir.x,
		y: v3.pos.y + v3.dir.y,
		z: v3.pos.z + v3.dir.z,
	}

	//for part 1 we only consider X and Y
	//we are not checking for 0 divide here for now assuming that this case doesn't appear
	M := (B.y - A.y) / (B.x - A.x)

	// y = mx+p
	// y-mx = p
	P := A.y - M*A.x

	return M, P
}

func findIntersectBetweenLines(l1M, l1P, l2M, l2P float64) (float64, float64) {
	/*
	 y = l1M * x + l1P
	 y = l2M * x + l2P
	 (l1M * x) + l1P = (l2M * x) + l2P
	 (l1M * x) - (l2M * x) = l2P - l1P
	 x(L1M - l2M) = l2P - l1P
	 x = (l2P - l1P) / (L1M - l2M)
	*/

	//we are not checking for 0 divide here for now assuming that this case doesn't appear
	x := (l2P - l1P) / (l1M - l2M)
	y := l1M*x + l1P

	return x, y
}

func checkPositionInTestArea(lowBound, upperBound float64, vA, vB Vector3) bool {
	aM, aP := findLineEquation(vA)
	bM, bP := findLineEquation(vB)

	intersectX, intersectY := findIntersectBetweenLines(aM, aP, bM, bP)

	//get direction of movment then direction toward intersect
	vAA := vA.dir
	vAI := Point3{intersectX - vA.pos.x, intersectY - vA.pos.y, 0}
	vBB := vB.dir
	vBI := Point3{intersectX - vB.pos.x, intersectY - vB.pos.y, 0}

	//get dot product bewteen the 2 direction, positive mean toward future
	dotA := dot(vAA, vAI, 2)
	dotB := dot(vBB, vBI, 2)

	inThePastA := dotA < 0
	inThePastB := dotB < 0

	outOfBounds := intersectX < lowBound ||
		intersectX > upperBound ||
		intersectY < lowBound ||
		intersectY > upperBound

	result := !outOfBounds && !(inThePastA || inThePastB)

	return result
}

//_____________________________________________________________________________
//______________________________________PART 2_________________________________
//_____________________________________________________________________________
