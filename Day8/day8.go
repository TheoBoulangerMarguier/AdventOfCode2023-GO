/*
--- Day 8: Haunted Wasteland ---
You're still riding a camel across Desert Island when you spot a sandstorm
quickly approaching. When you turn to warn the Elf, she disappears before your
eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents
(your puzzle input) about how to navigate the desert. At least, you're pretty
sure that's what they are; one of the documents contains a list of left/right
instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the n
etwork. Perhaps if you have the camel follow the same instructions,
you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ.
You feel like AAA is where you are now, and you have to follow the
left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
Starting with AAA, you need to look up the next element based on the next
left/right instruction in your input. In this example, start with AAA and go
right (R) by choosing the right element of AAA, CCC. Then, L means to choose
the left element of CCC, ZZZ. By following the left/right instructions,
you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right
instructions, repeat the whole sequence of instructions as necessary:
RL really means RLRLRLRLRLRLRLRL... and so on. For example,
here is a situation that takes 6 steps to reach ZZZ:

LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
Starting at AAA, follow the left/right instructions.
How many steps are required to reach ZZZ?

--- Part Two ---
The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
Here, there are two starting nodes, 11A and 22A (because they both end with A).
As you follow each left/right instruction, use that instruction to simultaneously
navigate away from both nodes you're currently on. Repeat this process until
all of the nodes you're currently on end with Z. (If only some of the nodes
you're on end with Z, they act like any other node and you continue as normal.)
In this example, you would proceed as follows:

Step 0: You are at 11A and 22A.
Step 1: You choose all of the left paths, leading you to 11B and 22B.
Step 2: You choose all of the right paths, leading you to 11Z and 22C.
Step 3: You choose all of the left paths, leading you to 11B and 22Z.
Step 4: You choose all of the right paths, leading you to 11Z and 22B.
Step 5: You choose all of the left paths, leading you to 11B and 22C.
Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take
before you're only on nodes that end with Z?

*/

package Day8

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"log"
	"os"
	"regexp"
	"sync"
)

func Day8() [2]int {
	result := [2]int{
		d8p1(),
		d8p2(),
	}
	return result

}

type node struct {
	name     string
	lastChar string
	left     string
	right    string
}

func createTree(path string) ([]rune, map[string]node) {
	//step 1 open the file and scan it
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	instructions := []rune{}
	tree := map[string]node{}

	scanner := bufio.NewScanner(file)

	instructionFinihed := false
	for scanner.Scan() {
		if !instructionFinihed {
			if scanner.Text() == "" {
				instructionFinihed = true
			} else {
				instructions = append(instructions, []rune(scanner.Text())...)
			}
		} else {

			re := regexp.MustCompile(`\b[A-Za-z0-9]{3}\b`)
			names := re.FindAllString(scanner.Text(), -1)

			n := node{
				name:     names[0],
				lastChar: string([]rune(names[0])[2]),
				left:     names[1],
				right:    names[2],
			}

			tree[names[0]] = n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions, tree
}

func walkWithPatternUntilEqual(tree map[string]node, instructions []rune, a string, b string, useLast bool) int {
	stepCount := 0
	start := a
	end := b
	current := start

	for {
		for i := 0; i < len(instructions); i++ {
			if instructions[i] == 'L' {
				current = tree[current].left
			} else {
				current = tree[current].right
			}
			stepCount++

			if useLast {
				if tree[current].lastChar == end {
					return stepCount
				}
			} else {
				if current == end {
					return stepCount
				}
			}

		}
	}
}

func d8p1() int {
	instructions, tree := createTree("./Day8/Ressources/day8_input.txt")
	return walkWithPatternUntilEqual(tree, instructions, "AAA", "ZZZ", false)
}

func d8p2() int {
	instructions, tree := createTree("./Day8/Ressources/day8_input.txt")
	startingNodes := []string{}
	pathsStepCount := []int{}

	startChar := "A"
	endChar := "Z"

	//find all starting
	for k := range tree {
		if tree[k].lastChar == startChar {
			startingNodes = append(startingNodes, k)
		}
	}

	//get the step count to finding a Z of each starting node (multithreaded)
	var wg sync.WaitGroup
	wg.Add(len(startingNodes))

	for i := 0; i < len(startingNodes); i++ {
		go func(i int) {
			defer wg.Done()
			n := walkWithPatternUntilEqual(tree, instructions, startingNodes[i], endChar, true)
			pathsStepCount = append(pathsStepCount, n)
		}(i)
	}
	wg.Wait()

	//calculates the Least Common Multiple (LCM) of a list of numbers
	result := 1
	for _, num := range pathsStepCount {
		gcd := utils.LCM(result, num) // Calculate GCD of 'result' and 'num'
		result = (result / gcd) * num // Update 'result'
	}
	return result
}
