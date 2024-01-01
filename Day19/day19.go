/*
--- Day 19: Aplenty ---
The Elves of Gear Island are thankful for your help and send you on your way.
They even have a hang glider that someone stole from Desert Island;
since you're already going that direction, it would help them a lot
if you would use it to get down there and return it to them.

As you reach the bottom of the relentless avalanche of machine parts,
you discover that they're already forming a formidable heap.
Don't worry, though - a group of Elves is already here organizing the parts,
and they have a system.

To start, each part is rated in each of four categories:

x: Extremely cool looking
m: Musical (it makes a noise when you hit it)
a: Aerodynamic
s: Shiny

Then, each part is sent through a series of workflows that will ultimately accept
or reject the part.
Each workflow has a name and contains a list of rules; each rule specifies a
condition and where to send the part if the condition is true.
The first rule that matches the part being considered is applied immediately,
and the part moves on to the destination described by the rule.
(The last rule in each workflow has no condition and always applies if reached.)

Consider the workflow ex{x>10:one,m<20:two,a>30:R,A}.
This workflow is named ex and contains four rules.
If workflow ex were considering a specific part, it would perform the
following steps in order:

Rule "x>10:one": If the part's x is more than 10, send the part to the workflow named one.
Rule "m<20:two": Otherwise, if the part's m is less than 20, send the part to the workflow named two.
Rule "a>30:R": Otherwise, if the part's a is more than 30, the part is immediately rejected (R).
Rule "A": Otherwise, because no other rules matched the part, the part is immediately accepted (A).
If a part is sent to another workflow, it immediately switches to the start of that workflow instead and never returns.
If a part is accepted (sent to A) or rejected (sent to R), the part immediately stops any further processing.

The system works, but it's not keeping up with the torrent of weird metal shapes.
The Elves ask if you can help sort a few parts and give you the list of workflows and some part ratings (your puzzle input). For example:

px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}

The workflows are listed first, followed by a blank line, then the ratings of the parts the Elves would like you to sort.
All parts begin in the workflow named in. In this example, the five listed parts go through the following workflows:

{x=787,m=2655,a=1222,s=2876}: in -> qqz -> qs -> lnx -> A
{x=1679,m=44,a=2067,s=496}: in -> px -> rfg -> gd -> R
{x=2036,m=264,a=79,s=2244}: in -> qqz -> hdj -> pv -> A
{x=2461,m=1339,a=466,s=291}: in -> px -> qkq -> crn -> R
{x=2127,m=1623,a=2188,s=1013}: in -> px -> rfg -> A

Ultimately, three parts are accepted. Adding up the x, m, a, and s rating for each of the accepted parts gives 7540 for the part with x=787, 4623
for the part with x=2036, and 6951 for the part with x=2127. Adding all of the ratings for all of the accepted parts gives the sum total of 19114.

Sort through all of the parts you've been given;
what do you get if you add together all of the rating numbers for all of the parts that ultimately get accepted?
*/

package Day19

import (
	utils "AdventOfCode/Utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day19() [2]int {
	return [2]int{
		d19p1(),
		d19p2(),
	}
}

type Workflow struct {
	rules []Rule
}

type Rule struct {
	rating       string
	condition    bool // "<" false, ">" true
	threshold    int
	sendToAdress string
}

func loadData() (map[string]Workflow, []map[string]int) {
	file, err := os.Open("./Day19/Ressources/day19_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	workflows := map[string]Workflow{}
	parts := []map[string]int{}

	isWorkflow := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			isWorkflow = false
			continue
		}

		if isWorkflow {
			split := strings.Split(scanner.Text(), "{")
			name := split[0]
			split = strings.Split(split[1], "}")
			rulesStrings := strings.Split(split[0], ",")
			rules := []Rule{}
			for _, r := range rulesStrings {
				chars := []rune(r)

				ok, _ := utils.SliceContains(chars, ':')
				rating, sendToAdress := "", ""
				condition := false
				threshold := -1

				if ok { //has a re-assignment condition
					rating = string(chars[0]) //set letter
					if chars[1] == '>' {
						condition = true //set condition
					}
					index, _ := utils.SliceLastIndexOf(chars, ':')
					tString := string(chars[2:index])
					v, err := strconv.Atoi(tString)
					if err != nil {
						log.Fatal(err)
					}
					threshold = v                           //set threshold
					sendToAdress = strings.Split(r, ":")[1] //set send to

				} else { //has direct re-assignment
					sendToAdress = r
				}

				newRule := Rule{
					rating:       rating,
					condition:    condition,
					threshold:    threshold,
					sendToAdress: sendToAdress,
				}

				rules = append(rules, newRule)
			}

			newWorkflow := Workflow{
				rules: rules,
			}

			workflows[name] = newWorkflow

		} else {
			split := strings.Split(string([]rune(scanner.Text())[1:len(scanner.Text())-1]), ",")
			xString := strings.Split(split[0], "=")[1]
			mString := strings.Split(split[1], "=")[1]
			aString := strings.Split(split[2], "=")[1]
			sString := strings.Split(split[3], "=")[1]

			xInt, err0 := strconv.Atoi(xString)
			if err0 != nil {
				log.Fatal(err0)
			}

			mInt, err1 := strconv.Atoi(mString)
			if err1 != nil {
				log.Fatal(err1)
			}

			aInt, err2 := strconv.Atoi(aString)
			if err2 != nil {
				log.Fatal(err2)
			}

			sInt, err3 := strconv.Atoi(sString)
			if err3 != nil {
				log.Fatal(err3)
			}

			newPart := map[string]int{
				"x": xInt,
				"m": mInt,
				"a": aInt,
				"s": sInt,
			}
			parts = append(parts, newPart)

		}
	}

	return workflows, parts
}

func d19p1() int {
	workflows, parts := loadData()
	validParts := []map[string]int{}
	for _, p := range parts {
		result := processPart(workflows, "in", p)
		if result {
			validParts = append(validParts, p)
		}
	}
	fmt.Println(validParts)
	sum := 0
	for _, vp := range validParts {
		sum += vp["x"] + vp["m"] + vp["a"] + vp["s"]
	}
	return sum
}

func processPart(workflows map[string]Workflow, wID string, part map[string]int) bool {
	workflow := workflows[wID]
	for _, r := range workflow.rules {
		if r.threshold == -1 {
			if r.sendToAdress == "A" {
				//accept
				return true
			} else if r.sendToAdress == "R" {
				//reject
				return false
			} else {
				//send to other workflow
				return processPart(workflows, r.sendToAdress, part)
			}
		} else {
			if (r.condition && part[r.rating] > r.threshold) ||
				(!r.condition && part[r.rating] < r.threshold) {
				if r.sendToAdress == "A" {
					//accept
					return true
				} else if r.sendToAdress == "R" {
					//reject
					return false
				} else {
					//send to other workflow
					return processPart(workflows, r.sendToAdress, part)
				}
			}
		}
	}

	return true
}

func d19p2() int {
	return 0
}
